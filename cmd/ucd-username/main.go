package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aaronland/go-ucd-username"	
)

func main() {

	var spaces = flag.Bool("spaces", false, "Do not filter out whitespace during processing")
	var punct = flag.Bool("punct", false, "Do not filter out punctuation during processing")
	var debug = flag.Bool("debug", false, "Enable verbose logging during processing")

	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, "Command line tool for converting strings in to valid UCD usernames.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] string(N) string(N) string(N)\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "For example:\n\t%s captain 🧍 ✨ \n\taptainastandingpersonsarkles\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nValid options are:\n")
		flag.PrintDefaults()
	}

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
