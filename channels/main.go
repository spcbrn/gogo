package main

import (
	"fmt"
	"net/http"
	"time"
)

type fn func()

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

	// could use syntax `for l := range c...`
	// but can also imply with `for...` and <-c reference within the loop
	for {
		fmt.Println("checking", <-c)
		go checkLink(<-c, c, true)
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
			// delayFeedChannel(url, c, 2)
			sleepExec(func() { c <- url }, 2)
		}
		return
	}
	fmt.Println(url, " appears to be up.")
	if loopChan == true {
		// delayFeedChannel(url, c, 2)
		sleepExec(func() { c <- url }, 2)
	}
}

func sleepExec(cb fn, sec time.Duration) {
	time.Sleep(time.Second * sec)
	cb()
}

// func delayFeedChannel(url string, c chan string, sec time.Duration) {
// 	time.Sleep(time.Second * sec)
// 	c <- url
// }
