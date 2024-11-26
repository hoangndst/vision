SHELL = /bin/bash
PWD:=$(shell pwd)
OS := $(shell uname -s)

GOFORMATER			?= gofumpt
GOFORMATER_VERSION	?= v0.2.0
GOLINTER			?= golangci-lint
GOLINTER_VERSION	?= v1.56.2
COVER_FILE			?= coverage.out
SOURCE_PATHS		?= .

.DEFAULT_GOAL := help

help:  ## This help message :)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test:  ## Run the tests
	go test -gcflags="all=-l -N" -timeout=10m `go list $(SOURCE_PATHS)` ${TEST_FLAGS}

cover:  ## Generates coverage report
	go test -gcflags="all=-l -N" -timeout=10m `go list $(SOURCE_PATHS)` -coverprofile $(COVER_FILE) ${TEST_FLAGS}

cover-html:  ## Generates coverage report and displays it in the browser
	go tool cover -html=$(COVER_FILE)

format:  ## Format source code
	@which $(GOFORMATER) > /dev/null || (echo "Installing $(GOFORMATER)@$(GOFORMATER_VERSION) ..."; go install mvdan.cc/gofumpt@$(GOFORMATER_VERSION) && echo -e "Installation complete!\n")
	@for path in $(SOURCE_PATHS); do ${GOFORMATER} -l -w -e `echo $${path} | cut -b 3- | rev | cut -b 5- | rev`; done;

lint:  ## Lint, will not fix but sets exit code on error
	@which $(GOLINTER) > /dev/null || (echo "Installing $(GOLINTER)@$(GOLINTER_VERSION) ..."; go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINTER_VERSION) && echo -e "Installation complete!\n")
	$(GOLINTER) run --deadline=10m $(SOURCE_PATHS)

lint-fix:  ## Lint, will try to fix errors and modify code
	@which $(GOLINTER) > /dev/null || (echo "Installing $(GOLINTER)@$(GOLINTER_VERSION) ..."; go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINTER_VERSION) && echo -e "Installation complete!\n")
	$(GOLINTER) run --deadline=10m $(SOURCE_PATHS) --fix

build-local-windows:  ## Build vision for windows
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./_build/windows/bin/vision.exe \
		-ldflags="-s -w" -tags rpc .

gen-api-spec: ## Generate API Specification with OpenAPI format
	@which swag > /dev/null || (echo "Installing swag@v1.16.3 ..."; go install github.com/swaggo/swag/cmd/swag@v1.16.3 && echo "Installation complete!\n")
	# Generate API documentation with OpenAPI format
	-swag init --parseDependency --parseInternal --parseDepth 1 -g ./vision.go -o api/openapispec/ && echo "🎉 Done!" || (echo "💥 Fail!"; exit 1)
	# Format swagger comments
	-swag fmt -g ./**/*.go && echo "🎉 Done!" || (echo "💥 Failed!"; exit 1)

gen-api-doc: ## Generate API Documentation by API Specification
	@which swagger > /dev/null || (echo "Installing swagger@v0.30.5 ..."; go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5 && echo "Installation complete!\n")
	-swagger generate markdown -f ./api/openapispec/swagger.json --output=docs/api.md && echo "🎉 Done!" || (echo "💥 Fail!"; exit 1)

.PHONY: test cover cover-html format lint lint-fix gen-api-spec gen-api-doc
