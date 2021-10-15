package main

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func getHTML() (string, error) {
	url := "https://go.dev/blog/"

	response, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "Error getting blog html")
	}

	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "Error reading blog html")
	}
	return string(html), nil
}
