//assignment using channels

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func check(url2 string, c chan string) {
	url1 := "https://www.google.com/search?q="
	url := url1 + url2
	fmt.Println(url)
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is Down", url)
		s += fmt.Sprintf("Error is %v", err)
		c <- s
	} else {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			s := fmt.Sprintf("Error in Reading the response of %v", url)
			c <- s
		} else {
			file := strings.Split(url, "=")[1]
			file += ".txt"

			//			c <- file
			err := ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				s := fmt.Sprintf("Error in writing the file %v", err)
				c <- s
			}
			c <- string(bodyBytes)
		}

	}
	elapsed := time.Since(start)
	fmt.Println("Execution Time", elapsed)
}

func main() {

	urls := []string{"iphone", "samsung", "xiaomi", "redmi", "oppo"}
	c := make(chan string)
	for _, url := range urls {
		go check(url, c)
	}
	fmt.Println("No. of Goroutines is ", runtime.NumGoroutine())
	fmt.Println()
	for i := 0; i < len(urls); i++ {

		fmt.Println(<-c)

	}
	close(c)
}
