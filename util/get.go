package util

import (
	"io"
	"net/http"
	"io/ioutil"
)
func GetHTML(url string) (string, bool) {
	resp, err := http.Get(url)
	
	if err != nil {
	  return "", false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  return "", false
	}

	html := string(body)
	return html, true;
}

func GetBody(url string) (io.ReadCloser, bool) {
  resp, err := http.Get(url)
	
	if err != nil {
	  return nil, false
	}

	return resp.Body, true;
}