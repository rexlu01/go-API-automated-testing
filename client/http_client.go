package client

import (
	"io"
	"net/http"
)

func HTTPrequest(mothod string, url string, body io.Reader) (*http.Request, error) {

	reqest, err := http.NewRequest(mothod, url, body)

	return reqest, err

}
