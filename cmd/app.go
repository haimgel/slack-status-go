package cmd

import (
	"fmt"
	"github.com/haimgel/slack-status/internal/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const AppName = "slack-status"
const defaultConfigFileName = "config.yaml"

type StatusConfig struct {
	Name     string        `mapstructure:"name"`
	Emoji    string        `mapstructure:"emoji"`
	Text     string        `mapstructure:"text"`
	Duration time.Duration `mapstructure:"duration"`
}

type ApplicationConfig struct {
	Accounts []string        `mapstructure:"accounts"`
	Statuses []*StatusConfig `mapstructure:"statuses"`
}

type Application struct {
	cmd        *cobra.Command
	cfgFile    string
	accounts   []string
	config     ApplicationConfig
	exitStatus int
	exit       func(int)
}

func Execute(version string, exit func(int), args []string) {
	app := newApplication(version, args, exit)
	app.execute(exit)
}

func newApplication(version string, args []string, exit func(int)) *Application {
	app := &Application{exit: exit}

	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:     "slack-status",
		Version: version,
		Short:   "Set Slack status for multiple accounts",
		Long: `Set Slack status to pre-defined values from a command line,
for multiple accounts at once`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return app.loadConfig()
		},
	}
	rootCmd.SetArgs(args)
	rootCmd.PersistentFlags().StringVar(&app.cfgFile, "config", defaultConfigFilePath(), "Configuration file")
	rootCmd.PersistentFlags().StringSliceVarP(&app.accounts, "use", "u", nil, "Slack accounts to query/set")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setCmd := &cobra.Command{
		Use:   "set [status name]",
		Short: "Set the specified status to the selected Slack accounts",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.parseStatusAndGetCookie(args[0], app.set)
		},
	}

	getCmd := &cobra.Command{
		Use:   "get [status name]",
		Short: "Query whether a specific status is currently in effect",
		Long: `Query whether a specific status is currently in effect.
The exit status would be '0' if the specified status is currently active
and '1' if it's not.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.parseStatusAndGetCookie(args[0], app.get)
		},
	}

	clearCmd := &cobra.Command{
		Use:   "clear",
		Short: "Clears the current status",
		Long:  `Clear the current status.`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.clear()
		},
	}

	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(clearCmd)
	app.cmd = rootCmd
	return app
}

func (app *Application) execute(exit func(int)) {
	err := app.cmd.Execute()
	if err != nil {
		exit(2)
	} else {
		exit(app.exitStatus)
	}
}

func (app *Application) loadConfig() error {
	viper.SetConfigFile(app.cfgFile)
	err := viper.BindPFlags(app.cmd.Flags())
	if err != nil {
		return err
	}
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&app.config)
	if err != nil {
		return err
	}
	for _, acc := range app.accounts {
		if !slices.Contains(app.config.Accounts, acc) {
			return fmt.Errorf("account '%s' is not defined", acc)
		}
	}
	if len(app.accounts) == 0 {
		app.accounts = app.config.Accounts
	}
	return nil
}

func (app *Application) parseStatusAndGetCookie(status string, action func(status *StatusConfig, dCookie *http.Cookie) error) error {
	statusConfig, err := app.parseStatus(status)
	if err != nil {
		return err
	}
	app.cmd.SilenceUsage = true
	cookie, err := slack.GetDCookie()
	if err != nil {
		return err
	}
	return action(statusConfig, cookie)
}

func (app *Application) get(status *StatusConfig, dCookie *http.Cookie) error {
	if len(app.accounts) != 1 {
		return fmt.Errorf("need to specify exactly one account to query")
	}
	fmt.Printf("Getting status '%s' for account '%s': ", status.Name, app.accounts[0])
	token, err := slack.ScrapeToken(app.accounts[0], dCookie)
	if err != nil {
		return err
	}
	statusStr, err := slack.GetStatus(token, dCookie.Value)
	if err != nil {
		return err
	}
	if statusStr == status.Text {
		fmt.Printf("Active\n")
		app.exitStatus = 0
	} else {
		fmt.Printf("Inactive\n")
		app.exitStatus = 1
	}
	return nil
}

func (app *Application) set(status *StatusConfig, dCookie *http.Cookie) error {
	for _, acc := range app.config.Accounts {
		if slices.Contains(app.accounts, acc) {
			fmt.Printf("Setting status '%s' for account '%s': ", status.Name, acc)
			token, err := slack.ScrapeToken(acc, dCookie)
			if err == nil {
				err = slack.SetStatus(token, dCookie.Value, status.Text, status.Emoji, status.Duration)
			}
			if err == nil {
				fmt.Printf("OK\n")
			} else {
				fmt.Printf("Error (%s)\n", err)
				app.exitStatus = 2
			}
		}
	}
	return nil
}

func (app *Application) clear() error {
	app.cmd.SilenceUsage = true
	dCookie, err := slack.GetDCookie()
	if err != nil {
		return err
	}
	for _, acc := range app.config.Accounts {
		if slices.Contains(app.accounts, acc) {
			fmt.Printf("Clearing status for account '%s': ", acc)
			token, err := slack.ScrapeToken(acc, dCookie)
			if err == nil {
				err = slack.SetStatus(token, dCookie.Value, "", "", 0)
			}
			if err == nil {
				fmt.Printf("OK\n")
			} else {
				fmt.Printf("Error (%s)\n", err)
				app.exitStatus = 2
			}
		}
	}
	return nil
}

func (app *Application) parseStatus(status string) (*StatusConfig, error) {
	for _, config := range app.config.Statuses {
		if config.Name == status {
			return config, nil
		}
	}
	return nil, fmt.Errorf("status '%s' is not found in the configuration", status)
}

func defaultConfigFilePath() string {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return filepath.Join("/", defaultConfigFileName)
	}
	return filepath.Join(cfgDir, AppName, defaultConfigFileName)
}
