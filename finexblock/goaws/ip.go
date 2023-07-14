package goaws

import (
	"io"
	"net/http"
)

func OwnPrivateIP() (ip string, err error) {
	var response *http.Response
	var bytes []byte

	response, err = http.Get("http://169.254.169.254/latest/meta-data/local-ipv4")
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	bytes, err = io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}