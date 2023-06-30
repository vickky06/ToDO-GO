package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var con *sql.DB //smallcase so only this package has access

func ConnectDB() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "v0s035z"
		password = "postgresPassword@11"
		dbname   = "v0s035z"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DB connection failed", err)
		return nil, err
	}

	fmt.Println("Connected to the database!")
	con = db
	return db, nil
}
