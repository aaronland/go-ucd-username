CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src/github.com/thisisaaronland/go-ucd-username; then rm -rf src/github.com/thisisaaronland/go-ucd-username; fi
	mkdir -p src/github.com/thisisaaronland/go-ucd-username
	cp *.go src/github.com/thisisaaronland/go-ucd-username/
	cp -r http src/github.com/thisisaaronland/go-ucd-username/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:	rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/aaronland/go-ucd"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-sanitize"
	rm -rf src/github.com/cooperhewitt/go-ucd
	mv src/github.com/aaronland/go-ucd src/github.com/cooperhewitt/

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt cmd/*.go
	go fmt http/*.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/ucd-username cmd/ucd-username.go
	@GOPATH=$(GOPATH) go build -o bin/ucd-usernamed cmd/ucd-usernamed.go

wasm:   self
	@GOPATH=$(GOPATH) GOARCH=wasm GOOS=js go build -o www/ucd.wasm cmd/ucd-wasm.go

docker-build:
	docker build -t ucd-username .

docker-debug: docker-build
	docker run -it -p 6161:8080 -e HOST='0.0.0.0' ucd-username
