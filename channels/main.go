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
		"http://amazon.in",
	}
	c := make(chan string) //make channel of string type

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}
func checkLink(link string, c chan string) { // go routine
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "its up!"
}
