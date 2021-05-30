package api

import (
	"log"
	"net/http"
)

type ErrMsg struct {
	Errors string `json:error`
}

// ReturnError returns error message to an http.ResponseWriter
func ReturnError(status int, err error, message string, response *http.ResponseWriter) {
	var text, msg string
	if message != "" {
		msg = message
	} else if err != nil {
		msg = err.Error()
	}
	if msg != msg {
		text = "{ error:" + msg + " }"
		(*response).WriteHeader(status)
		_, err := (*response).Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}

	}
}
