package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printErrorAndExit(err error) {
	fmt.Printf("%+v\n", err)
	os.Exit(1)
}

func main() {
	tmpFilename := "/tmp/html.html"
	//html, err := getHTML()
	//if err != nil {
	//	printErrorAndExit(err)
	//}

	//err = ioutil.WriteFile(tmpFilename, []byte(html), 0600)
	//if err != nil {
	//	printErrorAndExit(err)
	//}

	htmlBytes, err := ioutil.ReadFile(tmpFilename)
	if err != nil {
		printErrorAndExit(err)
	}
	html := string(htmlBytes)
	blogEntries, err := parseHTML(html)
	if err != nil {
		printErrorAndExit(err)
	}

	feed, err := generateFeed(blogEntries)
	if err != nil {
		printErrorAndExit(err)
	}

	fmt.Println(feed)
}
