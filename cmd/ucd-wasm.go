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

		pretty := args[0].String()
		
		safe, err := username.Translate(pretty)

		if err != nil {
			log.Fatal(err)
		}

		// js.Global().Set("safe", safe)
		return safe
	})

	defer ucd_username.Release()
	
	js.Global().Set("username", ucd_username)

	c := make(chan struct{}, 0)

	log.Println("WASM Go Initialized")
	<-c
}