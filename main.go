package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Books []Book

var books Books

func main() {
	populateBooks()
	handleRequests()
}

func populateBooks() {
	books = Books{
		Book{Title: "Book1", Desc: "Book1 Description", Content: "Book1 Content"},
		Book{Title: "Book2", Desc: "Book2 Description", Content: "Book2 Content"},
	}
}

func handleRequests() {
	fmt.Println("starting server...")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", allBooks)
	router.HandleFunc("/books/{key}", particularBook)
	http.ListenAndServe(":8083", router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, ": homePage function called")

	fmt.Fprintf(w, "Welcome to home page")
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, ": allBooks function called")

	json.NewEncoder(w).Encode(books)

}

func particularBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, ": particularBook function called")

	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["key"])

	json.NewEncoder(w).Encode(books[key-1])
}
