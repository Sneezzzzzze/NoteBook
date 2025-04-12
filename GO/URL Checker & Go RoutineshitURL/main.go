package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	// make is used to create a map
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.reddit.com/",
		"https://www.stackoverflow.com/",
		"https://www.quora.com/",
		"https://www.wikipedia.org/",
		"https://www.youtube.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Printf("URL: %s, Status: %s\n", url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url, status}
}
