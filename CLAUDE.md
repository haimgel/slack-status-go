# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

slack-status-go is a CLI tool for setting Slack status across multiple accounts simultaneously. It scrapes Slack cookies from browsers to authenticate and set predefined statuses via the Slack API.

## Architecture

### Core Components

- **main.go**: Entry point that builds version info and delegates to cmd package
- **cmd/app.go**: CLI application logic using Cobra, handles configuration loading and command dispatch
- **internal/slack/credentials.go**: Browser cookie scraping and Slack token extraction
- **internal/slack/status.go**: Slack API interactions for getting/setting status

### Key Architecture Details

1. **Authentication Flow**: Uses browser cookie scraping (kooky library) to extract Slack's 'd' cookie, then scrapes API tokens from Slack's emoji customization page
2. **Configuration**: YAML-based config with predefined statuses and account lists, loaded via Viper
3. **Multi-account Support**: Iterates through configured accounts, scraping tokens individually for each team
4. **Command Structure**: Three main commands - `set`, `get`, `clear` - all operating on predefined status configurations

### Dependencies

- Uses patched version of slack-go/slack (`github.com/rusq/slack`) for XOXC token support
- Configuration managed via spf13/viper and spf13/cobra

## Development Commands

### Building
```bash
go build -o slack-status .
```

### Testing
```bash
go test ./...
```

### Code Generation
```bash
go generate ./...
```

### Dependencies
```bash
go mod tidy
```

### Release Build (using goreleaser)
```bash
goreleaser build --snapshot --clean
```

## Configuration

- Default config location: `~/Library/Application Support/slack-status/config.yaml` (macOS)
- Sample config provided in `config-sample.yaml`
- Configuration defines:
  - `accounts`: List of Slack team names (subdomain before .slack.com)
  - `statuses`: Predefined status configurations with name, emoji, text, and duration

## Key Implementation Notes

- Browser cookie access requires keychain permissions on macOS
- Token scraping depends on specific HTML structure of Slack's emoji page
- Uses fake Chrome user agent to ensure proper token response from Slack
- Supports Go duration notation for status expiration (e.g., "60m" for 60 minutes)