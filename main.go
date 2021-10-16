package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func printErrorAndExit(err error) {
	log.Fatalf("%+v", err)
	os.Exit(1)
}

func main() {
	http.HandleFunc("/", handleHTTP)

	printErrorAndExit(http.ListenAndServe("localhost:8081", nil))
}
