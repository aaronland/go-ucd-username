package http

import (
	"github.com/thisisaaronland/go-ucd-username"
	"github.com/whosonfirst/go-sanitize"
	gohttp "net/http"
	"strconv"
)

type UCDUsernameHandlerOptions struct {
	Debug            bool
	AllowSpaces      bool
	AllowPunctuation bool
}

func UCDUsernameHandler(opts UCDUsernameHandlerOptions) (gohttp.Handler, error) {

	username, err := ucd.NewUCDUsername()

	if err != nil {
		return nil, err
	}

	username.Debug = opts.Debug
	username.AllowSpaces = opts.AllowSpaces
	username.AllowPunctuation = opts.AllowPunctuation

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		query := req.URL.Query()
		raw := query.Get("username")

		if raw == "" {
			gohttp.Error(rsp, "Missing username", gohttp.StatusBadRequest)
			return
		}

		opts := sanitize.DefaultOptions()

		scrubbed, err := sanitize.SanitizeString(raw, opts)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		if scrubbed == "" {
			gohttp.Error(rsp, "Invalid username", gohttp.StatusBadRequest)
			return
		}

		safe, err := username.Translate(scrubbed)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		b := []byte(safe)

		rsp.Header().Set("Content-Type", "text/plain")
		rsp.Header().Set("Content-Length", strconv.Itoa(len(b)))

		rsp.Write(b)
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
