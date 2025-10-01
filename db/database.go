package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// package level variable that will hold the datbase connection pool
//  a pointer doesnt hold the data itself; it holds the memory addres where the data is stored
//  a pointer doesnt hold the data itselfl it holds the memoery address where the data is stored


var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./tools.db")

	if err != nil {
		log.Fatal("Error Opening the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error Connecting the database: %v", err)
	}

	log.Println("Database Connection established")

	createTable()
}

func createTable() {

	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// db.exec executes a query without returning any rows
	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error Creating todos table: %v", err)
	}

	log.Println("Table 'Todos' is Ready")

}
