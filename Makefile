ifndef GOARCH
	GOARCH=$(shell go env GOARCH)
endif

ifndef GOOS
	GOOS := $(shell go env GOOS)
endif

ifndef GOPATH
	GOPATH=$(shell go env GOPATH)
endif

.PHONY: vendor
vendor:
	GO111MODULE=on go mod vendor

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: build
build:
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build -o bin/go-assign -mod=vendor cmd/generator/main.go

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