# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-admin
BINARY_UNIX=$(BINARY_NAME)_unix

all: run

build:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f ./build/$(BINARY_NAME)
	rm -f ./build/$(BINARY_UNIX)

run:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v ./
	./build/$(BINARY_NAME)

restart:
	kill -INT $$(cat pid)
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v ./
	./build/$(BINARY_NAME)

deps:
	$(GOGET) github.com/kardianos/govendor
	govendor sync

cross:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

pages:
	hero -source="resources/pages"
	mv resources/pages/*.go template

assets:
	go-bindata -o assets.go resources/assets/...