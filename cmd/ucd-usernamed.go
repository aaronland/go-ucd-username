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
	"github.com/facebookgo/grace/gracehttp"
	"github.com/thisisaaronland/go-ucd-username"
	"github.com/whosonfirst/go-sanitize"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var host = flag.String("host", "localhost", "What host to bind ucd-usernamed to")
	var port = flag.Int("port", 8080, "What port to bind ucd-usernamed to")

	var spaces = flag.Bool("spaces", false, "Do not filter out whitespace during processing")
	var punct = flag.Bool("punct", false, "Do not filter out punctuation during processing")
	var debug = flag.Bool("debug", false, "Enable verbose logging during processing")

	flag.Parse()

	handler := func(rsp http.ResponseWriter, req *http.Request) {

		query := req.URL.Query()
		raw := query.Get("username")

		if raw == "" {
			http.Error(rsp, "Missing username", http.StatusBadRequest)
			return
		}

		opts := sanitize.DefaultOptions()

		scrubbed, err := sanitize.SanitizeString(raw, opts)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		if scrubbed == "" {
			http.Error(rsp, "Invalid username", http.StatusBadRequest)
			return
		}

		username, err := ucd.NewUCDUsername()

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		username.Debug = *debug
		username.AllowSpaces = *spaces
		username.AllowPunctuation = *punct

		safe, err := username.Translate(scrubbed)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		b := []byte(safe)

		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Header().Set("Content-Length", strconv.Itoa(len(b)))

		rsp.Write(b)
	}

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("listening on", address)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	err := gracehttp.Serve(&http.Server{Addr: address, Handler: mux})

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
