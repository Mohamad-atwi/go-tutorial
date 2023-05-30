package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type MyWriter struct{}

func (mw MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p)) // Print the incoming byte slice as a string
	return len(p), nil    // Return the number of bytes written and a nil error
}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	myWriter := MyWriter{}
	_, _ = io.Copy(myWriter, resp.Body)
}
