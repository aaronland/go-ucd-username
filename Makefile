GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w
CWD=$(shell pwd)

vulnup:
	go install golang.org/x/vuln/cmd/govulncheck@latest

vuln:
	govulncheck -show verbose ./...

cli:
	@make wasm
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ucd-username cmd/ucd-username/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ucd-username-server cmd/ucd-username-server/main.go

wasm:   
	GOARCH=wasm GOOS=js go build -mod vendor -o http/wasm/ucd.wasm cmd/ucd-wasm/main.go

docker-build:
	docker build -t ucd-username .

docker-debug: docker-build
	docker run -it -p 6161:8080 -e HOST='0.0.0.0' ucd-username
