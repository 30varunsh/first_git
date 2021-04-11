package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type tokopedia struct {
	Keyword   string
	Response  string
	Time_took time.Duration
}

func main() {
	http.HandleFunc("/", view)
	http.ListenAndServe(":8081", nil)

}

func view(w http.ResponseWriter, req *http.Request) {
	c := make(chan tokopedia)
	v := req.FormValue("keywords") //FormValue to retrieve value from form url
	urls := strings.Split(v, ",")
	for _, url := range urls {
		go func(url2 string) {
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

				fmt.Println(strings.Repeat("#", 30))
				end := time.Now().UnixNano() / 1000000

				cal := time.Duration(end - start)
				fmt.Println("Time laga", cal)
				c <- tokopedia{url2, s, cal}
			}

			print(c, url2)
		}(url)
	}
	print2()
}


	bodyByte, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	s:=string(bodyByte)
}
