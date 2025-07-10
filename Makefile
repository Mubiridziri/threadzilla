.PHONY: build
build: ## Build a version
	make swag
	go build -v ./cmd/threadzilla

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: dev
dev: ## Go Run
	env $(shell cat .env) go run cmd/threadzilla/main.go

.PHONY: test
test: mocks
	go test ./...

.PHONY: lint
lint: ## Go Lint
	golangci-lint run ./...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
