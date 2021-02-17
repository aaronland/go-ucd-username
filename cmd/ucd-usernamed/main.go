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
	"fmt"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-ucd-username/http"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	fs := flagset.NewFlagSet("ucd")

	host := fs.String("host", "localhost", "What host to bind ucd-usernamed to. This fs is DEPRECATED. Please use -server-uri instead.")
	port := fs.Int("port", 8080, "What port to bind ucd-usernamed to. This fs is DEPRECATED. Please use -server-uri instead.")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	spaces := fs.Bool("spaces", false, "Do not filter out whitespace during processing")
	punct := fs.Bool("punct", false, "Do not filter out punctuation during processing")
	debug := fs.Bool("debug", false, "Enable verbose logging during processing")

	flagset.Parse(fs)

	ctx := context.Background()

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "UCD", false)

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

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
