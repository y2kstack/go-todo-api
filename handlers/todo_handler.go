package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/y2kstack/go-todo-api/db"
	"github.com/y2kstack/go-todo-api/models"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//  decode new request body

	var newTodo models.Todo

	err := json.NewDecoder(r.Body).Decode(&newTodo)

	if err != nil {
		log.Printf("Error decoding Request Body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if newTodo.Title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	createdTodo, err := db.InsertTodo(newTodo)

	if err != nil {
		log.Printf("Error inserting todo into database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdTodo)

	if err != nil {
		log.Printf("Error Encoding Response: %v", err)
	}
}
