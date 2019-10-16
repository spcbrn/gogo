package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://youtube.com",
		"https://facebook.com",
		"https://golang.org",
	}

	for _, url := range sites {
		_, err := http.Get(url)
		if err != nil {
			fmt.Println(url, "may be down!")
			return
		}
		fmt.Println(url, "appears to be up.")
	}
}
