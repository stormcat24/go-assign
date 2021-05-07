ifndef GOARCH
	GOARCH=$(shell go env GOARCH)
endif

ifndef GOOS
	GOOS := $(shell go env GOOS)
endif

ifndef GOPATH
	GOPATH=$(shell go env GOPATH)
endif

ROOT_PACKAGE := github.com/stormcat24/go-assign
VERSION_PACKAGE := $(ROOT_PACKAGE)/pkg/version
LDFLAG_OS_ARCH := "$(VERSION_PACKAGE).osArch"
LDFLAG_GIT_COMMIT := "$(VERSION_PACKAGE).gitCommit"
LDFLAG_GIT_COMMIT_FULL := "$(VERSION_PACKAGE).gitCommitFull"
LDFLAG_BUILD_DATE := "$(VERSION_PACKAGE).buildDate"
LDFLAG_VERSION := "$(VERSION_PACKAGE).version"

.PHONY: vendor
vendor:
	GO111MODULE=on go mod vendor

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: build
build:
	$(eval GIT_COMMIT := $(shell git describe --tags --always))
	$(eval GIT_COMMIT_FULL := $(shell git rev-parse HEAD))
	$(eval BUILD_DATE := $(shell date '+%Y%m%d'))
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build -ldflags "-s -w -X $(LDFLAG_OS_ARCH)=$(GOOS)/$(GOARCH) -X $(LDFLAG_GIT_COMMIT)=$(GIT_COMMIT) -X $(LDFLAG_GIT_COMMIT_FULL)=$(GIT_COMMIT_FULL) -X $(LDFLAG_BUILD_DATE)=$(BUILD_DATE) -X $(LDFLAG_VERSION)=$(BUILD_DATE)-$(GIT_COMMIT)" \
			-o bin/go-assign -mod=vendor cmd/generator/main.go

.PHONY: debug
debug: build
	@rm -rf $(GOPATH)/bin/go-assign
	@cp -R $(GOPATH)/src/github.com/stormcat24/go-assign/bin/go-assign $(GOPATH)/bin/
	@make gen-example

.PHONY: gen-example
gen-example:
	@go generate ./example/...

.PHONY: test
test:
	GO111MODULE=on go test -v ./...

.PHONY: update-credits
update-credits:
	@go install github.com/Songmu/gocredits/cmd/gocredits@latest
	@gocredits . > CREDITS
