package ucd

import (
	"errors"
	goucd "github.com/cooperhewitt/go-ucd"
	"github.com/whosonfirst/go-sanitize"
	_ "log"
	"regexp"
	"strings"
	"unicode"
)

var re_allowed *regexp.Regexp

func init() {
	re_allowed = regexp.MustCompile(`[a-zA-Z0-9\-]`)
}

func Safe(raw string) (string, error) {

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

	for _, r := range safe {

		if unicode.IsSpace(r) {
			continue
		}

		char := string(r)

		if char == "." {
			continue
		}

		if re_allowed.MatchString(char) {
			bits = append(bits, char)
			continue
		}

		name := goucd.Name(char)

		chars, err := Safe(name.Name)

		if err != nil {
			return "", err
		}

		bits = append(bits, chars)
	}

	safe = strings.Join(bits, "")
	safe = strings.ToLower(safe)

	return safe, nil
}
