package main

import (
	"github.com/thisisaaronland/go-ucd-username"
	"log"
	"syscall/js"
)

var username *ucd.UCDUsername
var ucd_username js.Func

func init() {

	u, err := ucd.NewUCDUsername()

	if err != nil {
		log.Fatal(err)
	}

	username = u

	username.Debug = true
	username.AllowSpaces = false
	username.AllowPunctuation = false
}

func main() {

	ucd_username = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		raw := args[0].String()

		safe, err := username.Translate(raw)

		if err != nil {
			log.Printf("Failed to translate '%s' because: %s\n", raw, err)
			return nil
		}

		return safe
	})

	defer ucd_username.Release()

	js.Global().Set("username", ucd_username)

	c := make(chan struct{}, 0)

	log.Println("WASM Go Initialized")
	<-c
}
