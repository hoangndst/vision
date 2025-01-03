SHELL = /bin/bash
PWD:=$(shell pwd)


RUN_IN_DOCKER:=docker run -it --rm
RUN_IN_DOCKER+=-v ~/.ssh:/root/.ssh
RUN_IN_DOCKER+=-v ~/.gitconfig:/root/.gitconfig
RUN_IN_DOCKER+=-v ~/go/pkg/mod:/go/pkg/mod
RUN_IN_DOCKER+=-v ${PWD}:/root/vision

GOFORMATER			?= gofumpt
GOFORMATER_VERSION	?= v0.2.0
GOLINTER			?= golangci-lint
GOLINTER_VERSION	?= v1.63.4
COVER_FILE			?= coverage.out
SOURCE_PATHS		?= ./...
BUILD_PATH 			?= ./_build/bin
ATLAS 				?= $(BUILD_PATH)/altas

.DEFAULT_GOAL := help

CCRED=\033[0;31m
CCEND=\033[0m

help:  ## This help message :)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# If you encounter an error like "panic: permission denied" on MacOS,
# please visit https://github.com/eisenxp/macos-golink-wrapper to find the solution.
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
	$(GOLINTER) run $(SOURCE_PATHS)

lint-fix:  ## Lint, will try to fix errors and modify code
	@which $(GOLINTER) > /dev/null || (echo "Installing $(GOLINTER)@$(GOLINTER_VERSION) ..."; go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINTER_VERSION) && echo -e "Installation complete!\n")
	$(GOLINTER) run $(SOURCE_PATHS) --fix

doc:  ## Start the documentation server with godoc
	@which godoc > /dev/null || (echo "Installing godoc@latest ..."; go install golang.org/x/tools/cmd/godoc@latest && echo -e "Installation complete!\n")
	godoc -http=:6060

clean:  ## Clean build bundles
	# Delete old artifacts
	-rm -rf ./_build/bundles

build-all: build-local-darwin-all build-local-linux-all build-local-darwin-arm64-all build-local-windows-all ## Build all platforms (darwin, linux, windows)

build-local-darwin:  ## Build vision tool chain for macOS
	# Delete old artifacts
	-rm -rf ./_build/bundles/vision-darwin
	# Build vision
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./_build/bundles/vision-darwin/bin/vision \
		-ldflags="-s -w" -tags rpc .

build-local-darwin-all: build-local-darwin ## Build vision for macOS
	# Copy docs
	cp -r ./docs ./_build/bundles/vision-darwin/docs
	
	# README.md
	cp ./README.md ./_build/bundles/vision-darwin
	# Build tgz
	cd ./_build/bundles/vision-darwin && tar -zcvf ../vision-darwin.tgz .
	cd ./_build/bundles && go run ../../hack/md5file/main.go vision-darwin.tgz > vision-darwin.tgz.md5.txt

build-local-darwin-arm64: ## Build vision tool chain for macOS arm64
	# Delete old artifacts
	-rm -rf ./_build/bundles/vision-darwin-arm64
	mkdir -p ./_build/bundles/vision-darwin-arm64/bin
	mkdir -p ./_build/bundles/vision-darwin-arm64/kclvm/bin

	# Build vision
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 \
		go build -o ./_build/bundles/vision-darwin-arm64/bin/vision \
		-ldflags="-s -w" -tags rpc .

build-local-darwin-arm64-all: build-local-darwin-arm64 ## Build vision for macOS arm64
	# Copy docs
	cp -r ./docs ./_build/bundles/vision-darwin-arm64/docs

	# README.md
	cp ./README.md ./_build/bundles/vision-darwin-arm64
	# Build tgz
	cd ./_build/bundles/vision-darwin-arm64 && tar -zcvf ../vision-darwin-arm64.tgz .
	cd ./_build/bundles && go run ../../hack/md5file/main.go vision-darwin-arm64.tgz > vision-darwin-arm64.tgz.md5.txt

build-local-linux-in-docker: ## Build visionctl-linux in docker
	${RUN_IN_DOCKER} make build-local-linux

build-local-linux-all-in-docker: ## Build visionctl-linux with kcl and kclopenapi in docker
	${RUN_IN_DOCKER} make build-local-linux-all

build-local-linux:  ## Build vision tool chain for linux
	# Delete old artifacts
	-rm -rf ./_build/bundles/vision-linux
	mkdir -p ./_build/bundles/vision-linux/bin

	# Build vision
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./_build/bundles/vision-linux/bin/vision \
		-ldflags="-s -w" -tags rpc .

build-local-linux-arm64:  ## Build vision tool chain for linux arm64
	# Delete old artifacts
	-rm -rf ./_build/bundles/vision-linux-arm64
	mkdir -p ./_build/bundles/vision-linux-arm64/bin

	# Build vision
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 \
		go build -o ./_build/bundles/vision-linux-arm64/bin/vision \
		-ldflags="-s -w" -tags rpc .

build-local-linux-all: build-local-linux  ## Build vision for linux
	# Copy docs
	cp -r ./docs ./_build/bundles/vision-linux/docs

	# Copy README.md
	cp ./README.md ./_build/bundles/vision-linux

	# Build tgz
	cd ./_build/bundles/vision-linux && tar -zcvf ../vision-linux.tgz  .
	cd ./_build/bundles && go run ../../hack/md5file/main.go vision-linux.tgz > vision-linux.tgz.md5.txt

build-local-windows:  ## Build vision tool chain for windows
	# Delete old artifacts
	-rm -rf ./_build/bundles/vision-windows
	mkdir -p ./_build/bundles/vision-windows/bin
	mkdir -p ./_build/bundles/vision-windows/kclvm/bin
	

	# Build vision
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./_build/bundles/vision-windows/bin/vision.exe \
		-ldflags="-s -w" -tags rpc .

build-local-windows-all: build-local-windows  ## Build vision for windows
	# Copy docs
	cp -r ./docs ./_build/bundles/vision-windows/docs

	# Copy README.md
	cp ./README.md ./_build/bundles/vision-windows
	# Build zip
	cd ./_build/bundles/vision-windows && zip -r ../vision-windows.zip .
	cd ./_build/bundles && go run ../../hack/md5file/main.go vision-windows.zip > vision-windows.zip.md5.txt

build-image:  ## Build vision image
	make build-local-linux-all
	docker build -t vision/vision .

sh-in-docker:  ## Run a shell in the docker container of vision
	${RUN_IN_DOCKER} bash

e2e-test:
	# Run e2e test
	hack/run-e2e.sh $(OSTYPE)

gen-api-spec: ## Generate API Specification with OpenAPI format
	@which swag > /dev/null || (echo "Installing swag@v1.16.3 ..."; go install github.com/swaggo/swag/cmd/swag@v1.16.3 && echo "Installation complete!\n")
	# Generate API documentation with OpenAPI format
	-swag init --parseDependency --parseInternal --parseDepth 1 -g ./vision.go -o api/openapispec/ && echo "🎉 Done!" || (echo "💥 Fail!"; exit 1)
	# Format swagger comments
	-swag fmt -g pkg/**/*.go && echo "🎉 Done!" || (echo "💥 Failed!"; exit 1)

gen-api-doc: ## Generate API Documentation by API Specification
	@which swagger > /dev/null || (echo "Installing swagger@v0.30.5 ..."; go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5 && echo "Installation complete!\n")
	-swagger generate markdown -f ./api/openapispec/swagger.json --output=docs/api.md && echo "🎉 Done!" || (echo "💥 Fail!"; exit 1)

##@ Database Migration
ATLAS_INSTALL_SCRIPT ?= "https://release.ariga.io/atlas/atlas-linux-amd64-latest"

.PHONY: atlas
atlas: $(ATLAS) ## Install atlas tools for linux amd64 only
$(ATLAS): | $(LOCALBIN)
	mkdir -p $(BUILD_PATH)
	test -s $(ATLAS) || curl -sL $(ATLAS_INSTALL_SCRIPT) -o $(ATLAS) && chmod +x $(ATLAS)

env ?= dev
.PHONY: migration
migration: atlas ## Generate migration file
	$(ATLAS) migrate diff --env $(env)
	$(ATLAS) migrate status --env $(env)

.PHONY: migrate
migrate: atlas ## Migrate database
	$(ATLAS) migrate apply --env $(env)
	$(ATLAS) migrate status --env $(env)

.PHONY: migrate-test
migrate-test: atlas ## Migrate database dry-run
	$(ATLAS) migrate apply --env $(env) --dry-run

# down migrate with input is specific version id eg: make migrate-down version=20210923123456
.PHONY: migrate-down
migrate-down: atlas ## Migrate database down to specific version
	$(ATLAS) migrate down --env $(env) --to-version $(version)
	$(ATLAS) migrate status --env $(env)

.PHONY: migrate-hash
migrate-hash: atlas ## Re-hash migration files
	$(ATLAS) migrate hash --env gorm
	$(ATLAS) migrate status --env gorm

.PHONY: migrate-status
migrate-status: atlas ## Show database status
	$(ATLAS) migrate status --env $(env)

.PHONY: test cover cover-html format lint lint-fix doc build-changelog upload clean build-all build-image build-local-linux build-local-windows build-local-linux-all build-local-windows-all e2e-test gen-api-spec gen-api-doc migrate migrate-down migrate-hash migrate-status migrate-test migration
