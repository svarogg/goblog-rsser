package main

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

type blogEntry struct {
	link    string
	title   string
	summary string
	author  string
}

func parseHTML(html string) ([]*blogEntry, error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var allEntries []*blogEntry

	var currentEntry *blogEntry
	document.Find("#blogindex").Children().EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i%2 == 0 { // blogtitle
			currentEntry = &blogEntry{}

			a := s.Find("a")
			relativeLink, exists := a.Attr("href")
			if !exists {
				err = errors.Errorf("No href in blogtitle")
				return false
			}
			currentEntry.link = "http://go.dev" + relativeLink
			currentEntry.title = a.Text()
			currentEntry.author = s.Find(".author").Text()

		} else { // blogsummary
			currentEntry.summary = strings.TrimSpace(s.Text())

			allEntries = append(allEntries, currentEntry)
		}
		return true
	})
	if err != nil {
		return nil, err
	}

	return allEntries, nil
}
