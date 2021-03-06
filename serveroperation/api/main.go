//json.Decoder is used if you need to decode multiple values from a stream of data,
//github.com/gorilla/mux
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct //(model)
type Book struct {
	ID     string  `json:"id"`     //json field name is id
	Isbn   string  `json:"isbn"`   //json field name is id
	Title  string  `json:"title"`  //json field name is id
	Author *Author `json:"author"` //json field name is id

}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init Books var as a slice Book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params (slug variable)
	//mux.Var will return data in map [key , value]
	//map[id,1]

	fmt.Println(params)
	//loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			fmt.Println(json.NewEncoder(w).Encode(item))

			return
		}
	}
	fmt.Println(json.NewEncoder(w).Encode(&Book{}))

}

//Create New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10_00_000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

//Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

func main() { //view
	//Init Router
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	//Route Handlers/ Establish Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
