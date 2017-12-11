package main

/*
curl -s -i 'http://localhost:8081?username=mr.+%F0%9F%98%81'
HTTP/1.1 200 OK
Content-Length: 29
Content-Type: text/plain
Date: Sat, 08 Apr 2017 01:02:58 GMT

mrgrinningfacewithsmilingeyes
*/

import (
	"flag"
	"fmt"
	"github.com/thisisaaronland/go-ucd-username/http"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	var host = flag.String("host", "localhost", "What host to bind ucd-usernamed to")
	var port = flag.Int("port", 8080, "What port to bind ucd-usernamed to")

	var spaces = flag.Bool("spaces", false, "Do not filter out whitespace during processing")
	var punct = flag.Bool("punct", false, "Do not filter out punctuation during processing")
	var debug = flag.Bool("debug", false, "Enable verbose logging during processing")

	flag.Parse()

	opts := http.UCDUsernameHandlerOptions{
		Debug:            *debug,
		AllowSpaces:      *spaces,
		AllowPunctuation: *punct,
	}

	handler, err := http.UCDUsernameHandler(opts)

	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("listening on", address)

	mux := gohttp.NewServeMux()
	mux.Handle("/", handler)

	err = gohttp.ListenAndServe(address, mux)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
