.DEFAULT_GOAL := help
VERSION ?= $(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD || git rev-parse --short HEAD)
BUILD_NUMBER ?= $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell LC_TIME=en_US.utf8 date -u)
COMMIT_HASH = $(shell git rev-parse HEAD)
DIST_ROOT=dist

# Go test sum configuration
GOTESTSUM_FORMAT ?= testname
GOTESTSUM_JUNITFILE ?= report.xml
GOTESTSUM_JSONFILE ?= gotestsum.json

######################## DO NOT MODIFY BELOW##########################################
GOPATH = $(shell go env GOPATH)
LDFLAGS += -X "github.com/pfltdv/izmir/model.Version=$(VERSION)"
LDFLAGS += -X "github.com/pfltdv/izmir/model.BuildNumber=$(BUILD_NUMBER)"
LDFLAGS += -X "github.com/pfltdv/izmir/model.BuildDate=$(BUILD_DATE)"
LDFLAGS += -X "github.com/pfltdv/izmir/model.CommitHash=$(COMMIT_HASH)"

export GOTESTSUM_FORMAT
export GOTESTSUM_JUNITFILE
export GOTESTSUM_JSONFILE
# MAIN TARGETS
.PHONY: checkstyle 
checkstyle: checkstyle/lint checkstyle/vet ## Runs style/lint checks

.PHONY: tools
tools: tools/lint tools/gotestsum ## Installs required tools

.PHONY: build
build: build/darwin/amd64 build/linux/amd64 build/windows/amd64 ## Compiles and generates the izmir binary
	
.PHONY: test
test: tools/gotestsum ## Executes tests
	GOMAXPROCS=1 $(GOPATH)/bin/gotestsum

.PHONY: run
run: ## Run izmir
	go run main.go $(ARGS)

# HELPERS
.PHONY: build/darwin/amd64
build/darwin/amd64:
	GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags '$(LDFLAGS)' -o dist/amd64/darwin/izmir .

.PHONY: build/linux/amd64
build/linux/amd64:
	GOOS=linux GOARCH=amd64 go build -trimpath -ldflags '$(LDFLAGS)' -o dist/amd64/linux/izmir .

.PHONY: build/windows/amd64
build/windows/amd64:
	GOOS=windows GOARCH=amd64 go build -trimpath -ldflags '$(LDFLAGS)' -o dist/amd64/windows/izmir.exe .

.PHONY: checkstyle/lint
checkstyle/lint: tools/lint ## Runs lint checks
	$(GOPATH)/bin/golangci-lint run ./...

.PHONY: checkstyle/vet
checkstyle/vet:  ## Runs vet checks
	go vet ./...

.PHONY: tools/lint
tools/lint: ## Installs golangci-lint tool
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: tools/gotestsum
tools/gotestsum:  ## Installs gotestsum tool
	go install gotest.tools/gotestsum@latest

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' ./Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	