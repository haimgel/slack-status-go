package config

import (
	"os"
	"path/filepath"
	"time"
)

const AppName = "slack-status"

type StatusConfig struct {
	Name     string        `mapstructure:"name"`
	Emoji    string        `mapstructure:"emoji"`
	Text     string        `mapstructure:"text"`
	Duration time.Duration `mapstructure:"duration"`
}

type ApplicationConfig struct {
	Accounts []string       `mapstructure:"accounts"`
	Statuses []StatusConfig `mapstructure:"statuses"`
}

func defaultConfigHome() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(cfgDir, 0750)
	return filepath.Join(cfgDir, AppName), err
}
