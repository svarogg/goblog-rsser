package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	feed, err := getFeed()
	if err != nil {
		printErrorAndExit(err)
	}

	_, err = w.Write([]byte(feed))
	if err != nil {
		err = errors.Wrap(err, "Error writing feed to response writer")
		printErrorAndExit(err)
	}
}

func getFeed() (string, error) {
	if isCacheHot() {
		fmt.Println("Cache is hot")
		return getCache(), nil
	}

	fmt.Println("Cache is cold... fetching feed")

	feed, err := fetchAndParseFeed()
	if err != nil {
		return "", err
	}
	setCache(feed)

	return feed, nil
}

func fetchAndParseFeed() (string, error) {
	html, err := getHTML()
	if err != nil {
		printErrorAndExit(err)
	}

	blogEntries, err := parseHTML(html)
	if err != nil {
		return "", err
	}

	feed, err := generateFeed(blogEntries)
	if err != nil {
		return "", err
	}
	return feed, nil
}
