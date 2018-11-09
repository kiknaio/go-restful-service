package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

type Book struct {
	ID		int	`json:id`
	Title	string	`json:title`
	Author	string	`json:author`
	Year	int	`json:year`
}

var books []Book

const (
	host     = "localhost"
	port     = 8080
	user     = "root"
	password = "toor"
	dbname   = "test"
)

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main()  {
	db, err := sql.Open("mysql", "root:toor@/goDB?charset=utf8")

	logFatal(err)

	defer db.Close()

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	logFatal(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	logFatal(err)

	id, err := res.LastInsertId()
	logFatal(err)

	fmt.Println(id)

	logFatal(err)

	//log.Println(pgUrl)
	//
	//router := mux.NewRouter()
	//
	//books = append(books,
	//	Book{ID:1, Title: "Golang pointer", Author: "Mr. Golang", Year: 2010},
	//	Book{ID:2, Title: "Mastering Go", Author: "Mr. Golang", Year: 2017},
	//	Book{ID:3, Title: "Golang for dummies", Author: "Dr.Yan Drewer", Year: 2015},
	//	Book{ID:4, Title: "Go familty", Author: "Sam Jackson", Year: 2013}, )
	//
	//
	//router.HandleFunc("/books", getBooks).Methods("GET")
	//router.HandleFunc("/books/{id}", getBook).Methods("GET")
	//router.HandleFunc("/books", addBook).Methods("POST")
	//router.HandleFunc("/books", updateBook).Methods("PUT")
	//router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	//
	//log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func addBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func removeBook(w http.ResponseWriter, r *http.Request) {
}
