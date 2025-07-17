# Threadzilla

<img src="resources/logo.png" width="300" height="300" alt="Threadzilla Logo" />

> Threadzilla is a Slack bot that generates threads based on the prompt provided by the user.

[![Go](https://github.com/Mubiridziri/threadzilla/actions/workflows/go.yml/badge.svg)](https://github.com/Mubiridziri/threadzilla/actions/workflows/go.yml)

--- 

## Installation

### Install dependencies

```bash
$ go mod download
```

### Make your .env file

### Environment variables

| Variable                     | Default | Description                                                                                                                                                                                                               |
|------------------------------|---------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DAEMON_MODE                  | false   | Should the bot be run in daemon mode? If true, the bot will run in a loop checking if it's time to send a message. If false, the bot will create a thread and exit. You can use this mode to run the bot with a cron job. |
| SENDING_HOUR_AT              | 10:09   | Time when the bot should send the message in 'HH:MM' format                                                                                                                                                               |
| SLACK_TOKEN                  |         | Slack token from your bot                                                                                                                                                                                                 |
| SLACK_CHANNEL                |         | Slack channel where the bot should send the message                                                                                                                                                                       |
| OPENAI_TOKEN                 |         | OpenAI token from your account                                                                                                                                                                                            |
| OPENAI_GENERATING_WITH_IMAGE | true    | Should the bot generate a message with an image?                                                                                                                                                                          |

For example:
```
DAEMON_MODE=true
SENDING_HOUR_AT=09:00
SLACK_TOKEN=token
SLACK_CHANNEL=CHANNEL_ID
OPENAI_TOKEN=token
OPENAI_GENERATING_WITH_IMAGE=true
```

### Run project

```bash
$ make dev
```

### Install golangci-lint for local use

```bash
$ brew install golangci-lint
```

After you can run:

```bash
$ make lint
```

## Other commands

```bash
$ make help
build                          Build a version
clean                          Remove temporary files
dev                            Go Run
lint                           Go Lint
```

## Project Structure

```text
📂cmd/
├─ 📂threadzilla
│  ├─ 📄main.go     // Main package of the application, containing minimal logic, only responsible for launching the application
📂internal/
├─ 📂app/           // Core application package. Dependencies are initialized here, main goroutines are started, and the web server is launched
├─ 📂config/        // Application configuration
├─ 📂service        // Business logic layer
├─ 📂utils          // Utility functions used across all layers of the application

```