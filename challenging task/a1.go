package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

type respData struct {
	Keyword  string `json:"keyword"`
	Response string `json:"response"`
}

func check(url2 string, c chan string, k chan string) {
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
		//var d1 respData
		//a := respData{Keyword: s}
		//res = append(res, a)
		//		d1.Keyword = s
		//byteData, err := json.Marshal(res)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(string(byteData))

	}
}
func main() {
	urls := []string{"iphone", "redmi"}
	// for _, url := range urls {
	// 	go check(url)
	// 	time.Sleep(time.Second * 15)

	// }
	//-------------------

	c := make(chan string)
	k := make(chan string)
	for _, url := range urls {
		go check(url, c, k)
	}

	for i := 0; i < len(urls); i++ { //receiving data from channel

		fmt.Println(<-k)
		fmt.Println(<-c)
	}

	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())
}
