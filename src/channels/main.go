package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// }
	// repeating routins
	// for {
	// 	go checkLink(<-c, c)
	// }
	// alternative syntax for repeating
	for l := range c {
		// function literal (anonymous function)
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!!!")
		c <- link
		return
	}

	fmt.Println(link, "is up and running")
	c <- link
}