package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8081", nil)
	//log.Fatal(err)

}
