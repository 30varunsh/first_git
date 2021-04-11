//form request: post,get
//post request: data send by request body
//get request: data send by url

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `
	<form method="get">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+v)
}
