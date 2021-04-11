//channels for concurrency of goroutines
//url:-   http://localhost:8081/?keywords=iphone,samsung,redmi,realme
//FormValue for fetching the keywords from url
//WriteString() method for writing data on the web browser

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type jsonData struct {
	Keyword   string `json:"keyword"`
	Response  string `json:"response"`
	Time_took string `json:"time_took"`
}

type allJsonData []jsonData

func check(url2 string, c chan string) {
	url1 := "https://www.google.com/search?q="
	url := url1 + url2
	fmt.Println(url)
	start := time.Now()
	resp, err := http.Get(url) //send get request
	if err != nil {
		// 	s := fmt.Sprintf("%s is Down", url)
		// 	s += fmt.Sprintf("Error is %v", err)
		// 	c <- s
		//
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//s := fmt.Sprintf("Error in Reading the response of %v", url)
			//c <- s
			fmt.Println(err)
		} else {
			file := strings.Split(url, "=")[1]
			file += ".txt"
			err := ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				//s := fmt.Sprintf("Error in writing the file %v", err)
				//c <- s
				fmt.Println(err)
			}
			c <- string(bodyBytes)

			elapsed := time.Since(start)
			fmt.Println("execution time", elapsed)
			fmt.Printf("execution time format is %T", elapsed)
			c <- string(elapsed)
		}

	}
}

func main() {
	http.HandleFunc("/", view)
	http.ListenAndServe(":8081", nil)
}
func view(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("keywords") //FormValue to retrieve value from form
	x := strings.Split(v, ",")
	io.WriteString(w, "Do my search : "+v) //used to write in the browser

	c := make(chan string)
	for _, url := range x {
		go check(url, c)
	}
	fmt.Println("No. of Goroutines is ", runtime.NumGoroutine())
	fmt.Println()
	for i := 0; i < len(x); i++ { //receiving data from channel
		fmt.Println(<-c)

		fmt.Println(<-c)

	}
	close(c)
}
