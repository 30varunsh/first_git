package main

import (
	"net/http"
	//"gowebexamples.com/http-server"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8081", srv)
}
