package db

import (
	"database/sql"

	"github.com/y2kstack/go-todo-api/models"
)

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

func GetTodoById(id int64) (models.Todo, error) {
	var todo models.Todo
	query := "SELECT id, title, completed, created_at FROM todos WHERE id = ?"

	err := DB.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Todo{}, err
		}
		return models.Todo{}, err
	}

	return todo, nil

}

func GetAllTodos() ([]models.Todo, error) {
	query := "SELECT id, title, completed, created_at FROM todos ORDER BY created_at DESC"

	// this will return multiple rows
	rows, err := DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo

		//  for each row scan the colum values into the fiels fot the todo struct
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}
	return todos, nil

}

func UpdateTodo(id int64, todo models.Todo) (models.Todo, error) {
	updateSQL := "UPDATE todos SET title = ?, completed = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"

	result, err := DB.Exec(updateSQL, todo.Title, todo.Completed, id)

	if err != nil {
		return models.Todo{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Todo{}, err
	}
	if rowsAffected == 0 {
		return models.Todo{}, sql.ErrNoRows
	}
	return GetTodoById(id)

}

func DeleteTodo(id int64) error {
	deleteSQL := "DELETE FROM todos WHERE id = ?"
	result, err := DB.Exec(deleteSQL, id)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
