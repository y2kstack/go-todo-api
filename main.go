package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/y2kstack/go-todo-api/db"
	"github.com/y2kstack/go-todo-api/handlers"
)

func main() {

	// initialize db
	db.InitDB()

	r := mux.NewRouter()

	// REGISTERS YOUR HANDLERS
	// HANDLES COLLECTION FOR ENDPOINT FOR GET ALL POST
	r.HandleFunc("/todos", handlers.TodoHandler).Methods("GET", "POST")

	r.HandleFunc("/todos/{id:[0-9]+}", handlers.SingleTodoHandler).Methods("GET", "PUT", "DELETE")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Go Todo API!!")
	})

	port := ":8080"
	fmt.Printf("Server is running on http:localhost%s\n", port)

	// http.HandleFunc("/todos", handlers.TodoHandler)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "welcome to the go todo API!")
	// })

	// port := ":8080"
	// fmt.Println("Server is running on localhost:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
