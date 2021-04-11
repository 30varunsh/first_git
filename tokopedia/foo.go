//array of structs
//marshal and encode in json
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Fname string `json:"first"`
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8081", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
		<html>
		<head>
		<title>FOO</title>
		</head>
		<body>
		You are at foo page
		</body>
		</html>`
	w.Write([]byte(s)) //print on the browser
}

func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := []person{
		person{
			"James",
			"Bond",
			[]string{"Suite", "Gun", "Sense of humor"},
		},
		person{
			"James",
			"Sharma",
			[]string{"Suite", "Gun", "Sense of humor"},
		},
	}
	byteBody, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(byteBody)
}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suite", "Glass", "Sense of Humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		fmt.Println(err)
	}
}
