cli: 	
	go build -mod vendor -o bin/ucd-username cmd/ucd-username/main.go
	go build -mod vendor -o bin/ucd-usernamed cmd/ucd-usernamed/main.go

wasm:   
	GOARCH=wasm GOOS=js go build -mod vendor -o www/ucd.wasm cmd/ucd-wasm/main.go

docker-build:
	docker build -t ucd-username .

docker-debug: docker-build
	docker run -it -p 6161:8080 -e HOST='0.0.0.0' ucd-username
