GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

cli:
	@make wasm
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/ucd-username cmd/ucd-username/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/ucd-username-server cmd/ucd-username-server/main.go

wasm:   
	GOARCH=wasm GOOS=js go build -mod $(GOMOD) -ldflags="-s -w" -o http/wasm/ucd.wasm cmd/ucd-wasm/main.go

docker-build:
	docker build -t ucd-username .

docker-debug: docker-build
	docker run -it -p 6161:8080 -e HOST='0.0.0.0' ucd-username
