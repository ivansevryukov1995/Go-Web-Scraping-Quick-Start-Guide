package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.packtpub.com/en-ru/product/mastering-kubernetes-9781804614754")
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	stringBody := string(data)

	re := regexp.MustCompile(`.*main-book-price.*\n.*(\$[0-9]*\.[0-9]{0,2})`)
	priceMatches := re.FindStringSubmatch(stringBody)

	fmt.Printf("Book Price: %s\n", priceMatches[1])
}
