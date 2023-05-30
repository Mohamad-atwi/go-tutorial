package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel) // Call checkLink function as a goroutine
	}

	for link := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second) // Sleep for 5 seconds
			checkLink(link, channel)    // Call checkLink function again after the delay
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link // Send link to the channel
		return
	}
	fmt.Println(link, "is up!")
	channel <- link // Send link to the channel
}
