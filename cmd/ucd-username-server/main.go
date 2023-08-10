package main

/*
curl -s -i 'http://localhost:8080/api?username=mr.+%F0%9F%98%81'
HTTP/1.1 200 OK
Content-Length: 29
Content-Type: text/plain
Date: Sat, 08 Apr 2017 01:02:58 GMT

mrgrinningfacewithsmilingeyes
*/

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-ucd-username/http/api"
	"github.com/aaronland/go-ucd-username/http/wasm"
	"github.com/sfomuseum/go-flags/flagset"	
)

func main() {

	fs := flagset.NewFlagSet("ucd")

	host := fs.String("host", "localhost", "What host to bind ucd-username-server to. This fs is DEPRECATED. Please use -server-uri instead.")
	port := fs.Int("port", 8080, "What port to bind ucd-username-server to. This fs is DEPRECATED. Please use -server-uri instead.")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	spaces := fs.Bool("spaces", false, "Do not filter out whitespace during processing")
	punct := fs.Bool("punct", false, "Do not filter out punctuation during processing")
	debug := fs.Bool("debug", false, "Enable verbose logging during processing")

	enable_api := fs.Bool("enable-api", true, "Enable the /api endpoint")
	enable_www := fs.Bool("enable-www", true, "Enable the / endpoint")

	fs.Usage = func() {

		fmt.Fprintf(os.Stderr, "HTTP server exposing the ucd-username functionality.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] \n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "For example:\n\t%s\n\t2021/02/17 08:55:00 Listening on http://localhost:8080\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nValid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	ctx := context.Background()

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "UCD", false)

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	if *server_uri == "" {
		*server_uri = fmt.Sprintf("http://%s:%d", *host, *port)
	}

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Unable to create server (%s), %v", *server_uri, err)
	}

	mux := http.NewServeMux()

	if *enable_www {

		www_handler, err := wasm.UCDUsernameWASMHandler()

		if err != nil {
			log.Fatalf("Failed to create WWW (wasm) handler, %v", err)
		}

		mux.Handle("/", www_handler)
	}

	if *enable_api {

		api_opts := api.UCDUsernameAPIHandlerOptions{
			Debug:            *debug,
			AllowSpaces:      *spaces,
			AllowPunctuation: *punct,
		}

		api_handler, err := api.UCDUsernameAPIHandler(api_opts)

		if err != nil {
			log.Fatalf("Failed to create API handler, %v", err)
		}

		mux.Handle("/api", api_handler)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

	os.Exit(0)
}
