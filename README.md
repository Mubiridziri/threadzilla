# Threadzilla

![Threadzilla Logo](resources/logo.png)

> Threadzilla is a Slack bot that generates threads based on the prompt provided by the user.

[![Go](https://github.com/Mubiridziri/threadzilla/actions/workflows/go.yml/badge.svg)](https://github.com/Mubiridziri/threadzilla/actions/workflows/go.yml)

--- 

## Installation

### Install dependencies

```bash
$ go mod download
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