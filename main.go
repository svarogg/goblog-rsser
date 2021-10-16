package main

import (
	"fmt"
	"net/http"
	"os"
)

func printErrorAndExit(err error) {
	fmt.Printf("%+v\n", err)
	os.Exit(1)
}

func main() {
	http.HandleFunc("/", handleHTTP)

	printErrorAndExit(http.ListenAndServe(":8081", nil))
}
