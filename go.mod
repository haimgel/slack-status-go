module github.com/haimgel/slack-status

go 1.18

require (
	github.com/slack-go/slack v0.11.3
	github.com/zellyn/kooky v0.0.0-20220429214451-68fd0a98b521
)

require (
	github.com/Velocidex/json v0.0.0-20220224052537-92f3c0326e5a // indirect
	github.com/Velocidex/ordereddict v0.0.0-20220411103415-79032cf99b1d // indirect
	github.com/Velocidex/yaml/v2 v2.2.8 // indirect
	github.com/bobesa/go-domain-util v0.0.0-20190911083921-4033b5f7dd89 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-ini/ini v1.66.4 // indirect
	github.com/go-sqlite/sqlite3 v0.0.0-20180313105335-53dd8e640ee7 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gonuts/binary v0.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/keybase/go-keychain v0.0.0-20220408132150-ad3b4a8fd4a7 // indirect
	github.com/zalando/go-keyring v0.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220408190544-5352b0902921 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.7 // indirect
	www.velocidex.com/golang/go-ese v0.1.0 // indirect
)

// Use a patched version because the official one does not support XOXC tokens: https://github.com/slack-go/slack/issues/996
replace github.com/slack-go/slack => github.com/rusq/slack v0.11.300
