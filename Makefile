VERSION := 1.0.0# Setting Release Version
PLATFORM := $(shell uname)# Getting Platform (Darwin, Linux)
MACHINE := $(shell uname -m)# Getting Machine type (x86_64, 386)

build_darwin:
	$(info Building for Darwin)
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/head-$(VERSION)-$(PLATFORM)-$(MACHINE) ./src/main.go

build_linux:
	$(info Building for Linux amd & arm)
	env GOOS=linux GOARCH=arm go build -o ./bin/head-$(VERSION)-$(PLATFORM)-arm ./src/main.go
	env GOOS=linux GOARCH=arm64 go build -o ./bin/head-$(VERSION)-$(PLATFORM)-arm ./src/main.go
	env GOOS=linux GOARCH=amd64 go build -o ./bin/head-$(VERSION)-$(PLATFORM)-amd64 ./src/main.go

build: build_linux build_darwin
	$(info Building for Linux & Darwin platforms)

clean:
	$(info Removing all binaries from)
	rm -rf ./bin/*

run:
	$(info You many need to build first if binary not found)
	./bin/head-$(VERSION)-$(PLATFORM)-$(MACHINE) --help
