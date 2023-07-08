package model

import (
	"database/sql"
	"fmt"
)

func CreateTodoTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS todo (
			id SERIAL PRIMARY KEY,
			description TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT current_timestamp
		)
	`)
	if err != nil {
		fmt.Println("Error creating TODO table:", err)
		return err
	}

	fmt.Println("TODO table created successfully")
	return nil
}
