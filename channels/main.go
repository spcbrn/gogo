package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://github.com",
		"https://youtube.com",
		"https://facebook.com",
		"https://golang.org",
		"https://christopherlemke.com",
		"https://suralink.com",
		"https://udemy.com",
	}

	c := make(chan string)

	for _, url := range sites {
		fmt.Println("checking", url)
		go checkLink(url, c, true)
	}

	for l := range c {
		fmt.Println("checking", l)
		go checkLink(l, c, true)
		// go func() {
		// 	time.Sleep(time.Second * 2)
		// 	checkLink(l, c, false)
		// }()
	}
}

func checkLink(url string, c chan string, loopChan bool) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, " may be down or unsecured!")
		if loopChan == true {
			delayFeedChannel(url, c, 2)
		}
		return
	}
	fmt.Println(url, " appears to be up.")
	if loopChan == true {
		delayFeedChannel(url, c, 2)
	}
}

func delayFeedChannel(url string, c chan string, sec time.Duration) {
	time.Sleep(time.Second * sec)
	c <- url
}
