package main

import (
	"fmt"
	"log"

	"net/http"

	book_controller "hello_go/bc"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server starting...")

	r := mux.NewRouter()
	book_controller.Init()

	r.HandleFunc("/books", book_controller.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", book_controller.GetBook).Methods("GET")
	r.HandleFunc("/books", book_controller.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", book_controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", book_controller.DeleteBook).Methods("DELETE")

	fmt.Println("Server working...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
