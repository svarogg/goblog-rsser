package main

import (
	"github.com/gorilla/feeds"
	_ "github.com/gorilla/feeds"
)

func generateFeed(blogEntries []*blogEntry) (string, error) {
	feed := &feeds.Feed{
		Title:       "go.dev blog",
		Description: "go.dev blog auto-generated RSS feed",
		Link:        &feeds.Link{Href: "http://cyka.com"},
	}

	for _, entry := range blogEntries {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       entry.title,
			Link:        &feeds.Link{Href: entry.link},
			Author:      &feeds.Author{Name: entry.author},
			Description: entry.summary,
		})
	}

	return feed.ToRss()
}
