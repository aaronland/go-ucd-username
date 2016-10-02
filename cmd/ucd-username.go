package main

import (
	"flag"
	"fmt"
	"github.com/thisisaaronland/go-ucd-username"
	"log"
)

func main() {

	flag.Parse()
	args := flag.Args()

	for _, user := range args {

		safe, err := ucd.Safe(user)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(user, safe)
	}

}
