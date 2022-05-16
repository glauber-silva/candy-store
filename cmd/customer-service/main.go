package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	url := "https://candystore.zimpler.net/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprint("status is not 200: %d", resp.StatusCode))
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(doc)
}
