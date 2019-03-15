package main

import (
	"github.com/thisisaaronland/go-ucd-username"
	"log"
	"syscall/js"
)

var username *ucd.UCDUsername

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

func ucd_username(raw []js.Value) {

	pretty := raw[0].String()
	safe, err := username.Translate(pretty)

	if err != nil {
		log.Println(err)
		return
	}
	
	js.Global().Set("username", safe)
}

func registerCallbacks() {
	js.Global().Set("username", js.NewCallback(ucd_username))
}

func main() {
	
	c := make(chan struct{}, 0)

	log.Println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
