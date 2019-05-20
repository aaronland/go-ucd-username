package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-ucd-username"
	"log"
	"os"
	"strings"
)

func main() {

	var spaces = flag.Bool("spaces", false, "Do not filter out whitespace during processing")
	var punct = flag.Bool("punct", false, "Do not filter out punctuation during processing")
	var debug = flag.Bool("debug", false, "Enable verbose logging during processing")

	flag.Parse()
	args := flag.Args()

	pretty := strings.Join(args, " ")

	uname, err := username.NewUCDUsername()

	if err != nil {
		log.Fatal(err)
	}

	uname.Debug = *debug
	uname.AllowSpaces = *spaces
	uname.AllowPunctuation = *punct

	safe, err := uname.Translate(pretty)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(safe)
	os.Exit(0)
}
