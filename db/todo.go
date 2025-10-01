package db

import "github.com/y2kstack/go-todo-api/models"

func InsertTodo(todo models.Todo) (models.Todo, error) {
	insertSQL := "INSERT INTO todos (title, completed) VALUES (?,?)"

	stmt, err := DB.Prepare(insertSQL)

	if err != nil {
		return models.Todo{}, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(todo.Title, todo.Completed)

	if err != nil {
		return models.Todo{}, err
	}

	newID, err := result.LastInsertId()

	if err != nil {
		return models.Todo{}, err
	}

	todo.ID = newID

	// query the full row to get all the fields
	var createdTodo models.Todo
	query := "SELECT id, title, completed, created_at FROM todos WHERE id = ?"

	err = DB.QueryRow(query, newID).Scan(
		&createdTodo.ID,
		&createdTodo.Title,
		&createdTodo.Completed,
		&createdTodo.CreatedAt,
	)

	if err != nil {
		return models.Todo{}, err
	}

	return createdTodo, nil

}
