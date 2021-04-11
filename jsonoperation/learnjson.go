//json.Marshal()

package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json:"title"`

	Author Author `json:"author"`
}

type Author struct {
	Name      string `jsong:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	fmt.Println("Hello World!!!")

	author := Author{Name: "Elliot forbes", Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}
	//fmt.Printf("%+v", book)

	//byteArray,err:=json.Marshal(book)
	byteArray, err := json.MarshalIndent(book, "", "   ")
	fmt.Printf("%T", byteArray) //type of json.Marshal is in []byte

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

}
