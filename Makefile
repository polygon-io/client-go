.DEFAULT_GOAL       := help
TARGET_MAX_CHAR_NUM := 20
VERSION             := ""

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: help fmt lint test rest-example ws-example test-coverage display-coverage release

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
	@go fmt ./...

## Lint with golangci-lint
lint:
	@echo Linting
	@golangci-lint run --no-config --issues-exit-code=0 --timeout=5m

## Run the tests
test:
	@echo Running tests
	@go test -race -v ./...

## Run the REST example
rest-example:
	@echo Running the REST example
	@go run rest/example/main.go

## Run the WebSocket example
ws-example:
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

## Publish a new release (usage: make release VERSION={VERSION_TAG})
release: fmt lint test
	@echo Tagging release with version '${VERSION}'
	@[[ "${VERSION}" == v* ]] || { echo "Must pass a version tag starting with 'v' (e.g. 'make release VERSION=v0.1.0')" ; exit 1; }
	@sed -i.bak '/const clientVersion/s/.*/const clientVersion = "${VERSION}"/' rest/client/client.go && rm rest/client/client.go.bak
	@git reset && git add -p rest/client/client.go
	@git checkout -b stage-${VERSION}
	@git commit -m "update client version tag to '${VERSION}'"
	@echo Creating and merging a PR
	@gh pr create --fill && gh pr merge --admin --squash --delete-branch
	@echo Publishing a release
	@gh release create ${VERSION}
