package ucd

import (
	"errors"
	goucd "github.com/cooperhewitt/go-ucd"
	"github.com/whosonfirst/go-sanitize"
	"log"
	"regexp"
	"strings"
	"unicode"
)

var re_allowed *regexp.Regexp
var debug bool

func init() {
	re_allowed = regexp.MustCompile(`[a-zA-Z0-9\-]`)
	debug = true
}

func Username(raw string) (string, error) {

	if debug {
		log.Println("PARSE", raw)
	}

	opts := sanitize.DefaultOptions()

	safe, err := sanitize.SanitizeString(raw, opts)

	if err != nil {
		return "", err
	}

	safe = strings.Trim(safe, " ")

	if safe == "" {
		return "", errors.New("Insufficient username")
	}

	bits := make([]string, 0)

	for i, r := range safe {

		if debug {
			log.Printf("RUNE %d %#U\n", i, r)
		}

		if unicode.IsSpace(r) {

			if debug {
				log.Printf("RUNE %d %#U is space\tSKIPPING\n", i, r)
			}

			continue
		}

		if unicode.IsPunct(r) {

			if debug {
				log.Printf("RUNE %d %#U is punctuation\tSKIPPING\n", i, r)
			}

			continue
		}

		char := string(r)

		if re_allowed.MatchString(char) {
			bits = append(bits, char)
			continue
		}

		if debug {
			log.Printf("RUNE %d %#U is not whitelisted\tPROCESSING\n", i, r)
		}

		name := goucd.Name(char)

		if name.Name == "" {
			return "", errors.New("Totally crazy-pants character!")
		}

		if debug {
			log.Printf("RUNE %d %#U return string '%s'\tPROCESSING\n", i, r, name.Name)
		}

		for j, r := range name.Name {

			log.Printf("RUNE %d:%d %#U\n", i, j, r)

			char = string(r)

			if !re_allowed.MatchString(char) {
				continue
			}

			bits = append(bits, char)
		}

	}

	if len(bits) == 0 {
		return "", errors.New("Nothing left to make a username with")
	}

	safe = strings.Join(bits, "")
	safe = strings.ToLower(safe)

	return safe, nil
}
