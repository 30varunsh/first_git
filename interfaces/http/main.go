package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// 	bs := make([]byte, 99999) //99999 empty byte slices
	// 	resp.Body.Read(bs)//Read method is using Body to get data
	// 	fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	//io.Copy(os.Stdout, resp.Body)
}
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
