package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
	"strings"
)

type respData struct {
	Keyword string `json:"keyword"`
}

func check(url2 string){
	url1 := "https://www.google.com/search?q="
	url := url1 + url2
	fmt.Println(url)
	resp, err := http.Get(url)
	var res []respData
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		s := string(body)


		fmt.Println(strings.Repeat("#",30))
		//var d1 respData
		a:=respData{Keyword:s}
		res = append(res,a)
//		d1.Keyword = s
		byteData, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(byteData))

	}
}
func main() {
	urls := []string{"iphone","redmi"}
	for _, url := range urls {
		go check(url)
		time.Sleep(time.Second * 15)
	}
	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())
}
