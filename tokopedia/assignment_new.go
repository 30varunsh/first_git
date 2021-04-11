//channels for concurrency of goroutines
//url:-   http://localhost:8081/?keywords=iphone,samsung,redmi,realme
//FormValue for fetching the keywords from url
//WriteString() method for writing data on the web browser

package main

import (
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

func check(url2 string, c chan respData) {

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

		//fmt.Println(strings.Repeat("#", 30))
		end := time.Now().UnixNano() / 1000000

		cal := time.Duration(end - start)
		//fmt.Println("Time_Took", cal)

		c <- respData{url2, s, cal}

	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		v := req.FormValue("keywords") //FormValue to retrieve value from form url
		urls := strings.Split(v, ",")
		c := make(chan respData)
		for _, url := range urls {
			go check(url, c)
		}

		// var resp []respData

		// for i := 0; i < len(urls); i++ { //receiving data from channel into struct data structure
		// 	a := respData{
		// 		Keyword: <-k, Response: <-c, TimeDuration: <-t,
		// 	}
		// 	resp = append(resp, a) //append data from struct into array of struct

		// }

		for i := 0; i < len(urls); i++ {
			fmt.Println(c)
		}
		// bodyByte, err := json.Marshal(resp)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(string(bodyByte))
		// io.WriteString(w, string(bodyByte))
		fmt.Println("No. of Goroutines :", runtime.NumGoroutine())

		close(c)

	})

	http.ListenAndServe(":8081", nil)
}
