package main

import (
	"fmt"
	"net/http"
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
		// call checkLink function here, maybe you should use some keyword before calling it?
	}

	for link := range channel {
		// your fancy function here
	}
}

func checkLink() {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// something should probably go here
		return
	}
	fmt.Println(link, "is up!")
	// something should probably go here
}
