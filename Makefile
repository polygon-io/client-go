.DEFAULT_GOAL       := help
TARGET_MAX_CHAR_NUM := 20

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: help fmt lint test example test-coverage display-coverage

## Show help
help:
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Format
fmt:
	@echo Formatting
	@go fmt .

## Lint with golangci-lint
lint:
	@echo Linting
	@golangci-lint run --no-config --issues-exit-code=0 --timeout=5m

## Run the tests
test:
	@echo Running tests
	@go test -race -v ./...

## Run the WebSocket example
example:
	@echo Running the WebSocket example
	@go run websocket/example/main.go

## Run the tests with coverage
test-coverage:
	@echo Running tests with coverage
	@go test ./... -short -coverprofile=cover.out -covermode=atomic -coverpkg=./...

## Display test coverage
display-coverage:
	@echo Displaying test coverage
	@go tool cover -html=cover.out
