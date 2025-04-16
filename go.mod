module github.com/haimgel/slack-status

go 1.24.2

require (
	github.com/browserutils/kooky v0.2.3
	github.com/slack-go/slack v0.16.0
	github.com/spf13/cobra v1.9.1
	github.com/spf13/viper v1.20.1
)

require (
	github.com/Velocidex/json v0.0.0-20220224052537-92f3c0326e5a // indirect
	github.com/Velocidex/ordereddict v0.0.0-20230909174157-2aa49cc5d11d // indirect
	github.com/Velocidex/yaml/v2 v2.2.8 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-sqlite/sqlite3 v0.0.0-20180313105335-53dd8e640ee7 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gonuts/binary v0.2.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/keybase/go-keychain v0.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/sagikazarmark/locafero v0.9.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/zalando/go-keyring v0.2.6 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	www.velocidex.com/golang/go-ese v0.2.0 // indirect
)

// Use a patched version because the official one does not support XOXC tokens: https://github.com/slack-go/slack/issues/996
replace github.com/slack-go/slack => github.com/rusq/slack v0.12.0

// Fix for Chrome 130 support: https://github.com/browserutils/kooky/issues/89
replace github.com/browserutils/kooky => github.com/DP19/kooky v0.2.3
