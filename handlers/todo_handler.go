package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/y2kstack/go-todo-api/db"
	"github.com/y2kstack/go-todo-api/models"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createdTodo(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func SingleTodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodo(w, r)
	case http.MethodPut:
		UpdateTodo(w, r)
	case http.MethodDelete:
		deleteTodo(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// get variables from the url
	// mux.vars returns a map of the route variables from the request URL
	fmt.Println("UpdateTodo called")
	fmt.Println("UpdateTodo called")
	fmt.Println("UpdateTodo called")
	fmt.Println("UpdateTodo called")
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		http.Error(w, "Id not found in the URL", http.StatusBadRequest)
		return
	}
	// ceonvert the id from the string to integer

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid Id Format", http.StatusBadRequest)
		return
	}
	// decode the request body
	var updatedTodo models.Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		log.Printf("Error decoding Request Body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if updatedTodo.Title == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}
	// call the database function to update the todo
	updatedTodo, err = db.UpdateTodo(id, updatedTodo)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo Not Found", http.StatusNotFound)
			return
		}
		log.Printf("Error updating todo in database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedTodo)
	if err != nil {
		log.Printf("Error Encoding Response: %v", err)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := db.GetAllTodos()
	if err != nil {
		log.Printf("Error Getting todos From database: %v", err)
		http.Error(w, "internal Server error", http.StatusInternalServerError)
		return
	}

	if todos == nil {
		todos = []models.Todo{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}

func createdTodo(w http.ResponseWriter, r *http.Request) {
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

func getTodo(w http.ResponseWriter, r *http.Request) {
	// get variables from the url
	// mux.vars returns a map of the route variables from the request URL

	// vars := mux.Vars(r)
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		http.Error(w, "Id not found in the URL", http.StatusBadRequest)
		return
	}

	// ceonvert the id from the string to integer

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid Id Format", http.StatusBadRequest)
	}

	// call the database function

	todo, err := db.GetTodoById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo Not Found", http.StatusNotFound)
			return
		}

		log.Printf("Error Getting tody by iD: %v", err)
		http.Error(w, "internal Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// get variables from the url
	// mux.vars returns a map of the route variables from the request URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Id not found in the URL", http.StatusBadRequest)
		return
	}
	// ceonvert the id from the string to integer
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid Id Format", http.StatusBadRequest)
		return
	}
	// call the database function to delete the todo
	err = db.DeleteTodo(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo Not Found", http.StatusNotFound)
			return
		}
		log.Printf("Error deleting todo from database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}
	// 204 no content is the standard response for successful delete operation
	w.WriteHeader(http.StatusNoContent)

}
