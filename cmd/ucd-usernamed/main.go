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
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-ucd-username/http"	
	"github.com/aaronland/go-http-server"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	host := flag.String("host", "localhost", "What host to bind ucd-usernamed to. This flag is DEPRECATED. Please use -server-uri instead.")
	port := flag.Int("port", 8080, "What port to bind ucd-usernamed to. This flag is DEPRECATED. Please use -server-uri instead.")

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	
	spaces := flag.Bool("spaces", false, "Do not filter out whitespace during processing")
	punct := flag.Bool("punct", false, "Do not filter out punctuation during processing")
	debug := flag.Bool("debug", false, "Enable verbose logging during processing")

	flag.Parse()

	if *server_uri == "" {
		*server_uri = fmt.Sprintf("http://%s:%d", *host, *port)
	}
	
	opts := http.UCDUsernameHandlerOptions{
		Debug:            *debug,
		AllowSpaces:      *spaces,
		AllowPunctuation: *punct,
	}

	ucd_handler, err := http.UCDUsernameHandler(opts)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Unable to create server (%s), %v", *server_uri, err)
	}

	mux := gohttp.NewServeMux()
	mux.Handle("/", ucd_handler)

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

	os.Exit(0)
}
