CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src/github.com/thisisaaronland/go-ucd-username; then rm -rf src/github.com/thisisaaronland/go-ucd-username; fi
	mkdir -p src/github.com/thisisaaronland/go-ucd-username
	cp *.go src/github.com/thisisaaronland/go-ucd-username/
	cp -r http src/github.com/thisisaaronland/go-ucd-username/
	cp -r vendor/src/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:	rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/cooperhewitt/go-ucd"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-sanitize"

vendor-deps: deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt cmd/*.go
	go fmt http/*.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/ucd-username cmd/ucd-username.go
	@GOPATH=$(GOPATH) go build -o bin/ucd-usernamed cmd/ucd-usernamed.go
