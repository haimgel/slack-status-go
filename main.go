package main

import (
	"fmt"
	"github.com/haimgel/slack-status/cmd"
	"os"
	"runtime"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
)

func buildVersion(version, commit, date string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	result = fmt.Sprintf("%s\ngoos: %s\ngoarch: %s", result, runtime.GOOS, runtime.GOARCH)
	return result
}

func main() {
	cmd.Execute(
		buildVersion(version, commit, date),
		os.Exit,
		os.Args[1:],
	)
}
