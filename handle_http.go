package main

import (
	"net/http"
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	html, err := getHTML()
	if err != nil {
		printErrorAndExit(err)
	}

	blogEntries, err := parseHTML(html)
	if err != nil {
		printErrorAndExit(err)
	}

	feed, err := generateFeed(blogEntries)
	if err != nil {
		printErrorAndExit(err)
	}

	_, err = w.Write([]byte(feed))
	if err != nil {
		printErrorAndExit(err)
	}
}
