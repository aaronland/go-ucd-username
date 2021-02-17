package wasm

import (
	"embed"
	"net/http"
)

//go:embed index.html
//go:embed wasm_exec.js
//go:embed ucd.wasm
var web_app embed.FS

func UCDUsernameWASMHandler() (http.Handler, error) {

	fs := http.FS(web_app)
	handler := http.FileServer(fs)

	return handler, nil
}
