module github.com/haimgel/slack-status

go 1.18

require (
	github.com/slack-go/slack v0.11.3
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.15.0
	github.com/zellyn/kooky v0.0.0-20221025221128-3e66d684c4db
	golang.org/x/exp v0.0.0-20230420155640-133eef4313cb
)

require (
	github.com/Velocidex/json v0.0.0-20220224052537-92f3c0326e5a // indirect
	github.com/Velocidex/ordereddict v0.0.0-20221110130714-6a7cb85851cd // indirect
	github.com/Velocidex/yaml/v2 v2.2.8 // indirect
	github.com/bobesa/go-domain-util v0.0.0-20190911083921-4033b5f7dd89 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-sqlite/sqlite3 v0.0.0-20180313105335-53dd8e640ee7 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gonuts/binary v0.2.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/keybase/go-keychain v0.0.0-20230307172405-3e4884637dd1 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/zalando/go-keyring v0.2.2 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	www.velocidex.com/golang/go-ese v0.1.0 // indirect
)

// Use a patched version because the official one does not support XOXC tokens: https://github.com/slack-go/slack/issues/996
replace github.com/slack-go/slack => github.com/rusq/slack v0.11.300
