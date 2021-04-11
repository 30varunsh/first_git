//Form Value
//important
//http://localhost:8081/?keywords=cow,buffalo
package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)
}
func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("keywords") //FormValue to retrieve value from form
	x := strings.Split(v, ",")
	_ = x
	io.WriteString(w, "Do my search : "+v) //used to write in the browser

}
