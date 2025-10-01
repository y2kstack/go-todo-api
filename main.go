package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/y2kstack/go-todo-api/db"
	"github.com/y2kstack/go-todo-api/handlers"
)

func main() {

	// initialize db
	db.InitDB()

	http.HandleFunc("/todos", handlers.TodoHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome to the go todo API!")
	})

	port := ":8080"
	fmt.Println("Server is running on localhost:", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
