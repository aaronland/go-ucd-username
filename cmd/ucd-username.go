package main

import (
	"flag"
	"fmt"
	"github.com/thisisaaronland/go-ucd-username"
	"log"
	"os"
	"strings"
)

func main() {

	flag.Parse()
	args := flag.Args()

	username := strings.Join(args, " ")

	safe, err := ucd.Username(username)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(safe)
	os.Exit(0)
}
