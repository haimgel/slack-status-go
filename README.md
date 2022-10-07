## SlackStatus - set Slack status for multiple accounts
[![Release](https://img.shields.io/github/release/haimgel/slack-status-go.svg?style=flat)](https://github.com/haimgel/slack-status-go/releases/latest)
[![Software license](https://img.shields.io/github/license/haimgel/slack-status-go.svg?style=flat)](/LICENSE)
[![Build status](https://img.shields.io/github/workflow/status/haimgel/slack-status-go/release.svg?style=flat)](https://github.com/haimgel/slack-status-go/actions?workflow=release)

This app allows to set Slack status to predefined values from a command line, for multiple accounts at once:
```shell
slack-status set lunch
```

This will set yourself away for an hour, with an hamburger emoji and "Lunch" status text.

*Note*: you could be prompted to give Keychain access to Chrome storage. This is one-time only request, and it's required
because `slack-status` scrapes Slack cookies from your browser, and uses that cookie to set status on your behalf.

## How to

```bash
# Add tap to your Homebrew
brew tap haimgel/tools

# Install it
brew install slack-status

# Configure (see below for details)
mkdir -p ~/Library/Application\ Support/slack-status
cp config-sample.yaml ~/Library/Application\ Support/slack-status/config.yaml
vi ~/Library/Application\ Support/slack-status/config.yaml
```

## Sample configuration file
```yaml
# List of accounts you wish to control. `slack-status` will scrape the cookie from your browser, so you don't need
# to configure anything besides the name of the team. The name of the team the part before .slack.com, e.g. if your
# Slack team is at https://team1.slack.com/, put just "team1" in this list.
accounts:
  - team1
  - team2
  - team3
# Pre-define statuses you want to set with this app. Be sure to quote the emoji name. Duration uses Go duration notation,
# e.g. 60m is 60 minutes.
statuses:
  - name: lunch
    emoji: ':hamburger:'
    text: Lunch
    duration: 60m
  - name: call
    emoji: ':phone:'
    text: In a call
    duration: 60m
  - name: brb
    emoji: ':brb:'
    text: BRB
    duration: 60m
```
