package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type respData struct {
	Keyword      string        `json:"keyword"`
	Response     string        `json:"response"`
	TimeDuration time.Duration `json:"time_took"`
}

func check(url2 string, c chan string, k chan string, t chan time.Duration) {

	start := time.Now().UnixNano() / 1000000
	url1 := "https://www.google.com/search?q="
	url := url1 + url2
	fmt.Println(url)
	resp, err := http.Get(url)
	//var res []respData
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		s := string(body)

		k <- url2
		c <- s
		fmt.Println(strings.Repeat("#", 30))
		end := time.Now().UnixNano() / 1000000

		cal := time.Duration(end - start)
		fmt.Println("Time laga", cal)
		t <- cal

	}
}
func main() {
	urls := []string{"iphone", "redmi"}
	c := make(chan string)
	k := make(chan string)
	t := make(chan time.Duration)
	for _, url := range urls {
		go check(url, c, k, t)
	}
	var resp []respData

	for i := 0; i < len(urls); i++ { //receiving data from channel into struct data structure
		a := respData{
			Keyword: <-k, Response: <-c, TimeDuration: <-t,
		}
		resp = append(resp, a) //append data from struct into array of struct

	}

	bodyByte, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bodyByte))
	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())
}
