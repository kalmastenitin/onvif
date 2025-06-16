package networking

import (
	"bytes"
	"log"
	"net/http"

	"github.com/juju/errors"
)

// SendSoap send soap message
func SendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		log.Printf("error in http post %v: %v", err.Error(), message)
		return resp, errors.Annotate(err, "Post")
	}

	return resp, nil
}
