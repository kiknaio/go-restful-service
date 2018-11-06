package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID		int	`json:id`
	Title	string	`json:title`
	Author	string	`json:author`
	Year	int	`json:year`
}

var books []Book

func main()  {
	router := mux.NewRouter()

	books = append(books,
		Book{ID:1, Title: "Golang pointer", Author: "Mr. Golang", Year: 2010},
		Book{ID:2, Title: "Mastering Go", Author: "Mr. Golang", Year: 2017},
		Book{ID:3, Title: "Golang for dummies", Author: "Dr.Yan Drewer", Year: 2015},
		Book{ID:4, Title: "Go familty", Author: "Sam Jackson", Year: 2013}, )


	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// LOGGING: Type
	//log.Println(reflect.TypeOf(params["id"]))

	// Convert String to Int
	i, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalln("Error while converting string to integer")
	}

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}
