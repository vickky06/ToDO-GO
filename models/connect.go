package model

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var con *sql.DB //smallcase so only this package has access

func Contains(arr []string, target string) bool {
	for _, str := range arr {
		if strings.Contains(str, target) {
			return true
		}
	}
	return false
}
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
	db1, err1 := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'")
	if err1 != nil {
		print("error while querying table")
	} else {
		var tableNames []string
		for db1.Next() {
			var tableName string
			err := db1.Scan(&tableName)
			if err != nil {
				fmt.Println("Error retrieving table name:", err)
				continue
			}
			tableNames = append(tableNames, tableName)
		}
		if len(tableNames) == 0 || !Contains(tableNames, "todo") {
			fmt.Println("No tables found. creating table Todo")
			err = CreateTodoTable(db)
			if err != nil {
				fmt.Println("Failed to create TODO table:", err)
				return nil, err
			}
		} else {
			fmt.Println("Table Names:", tableNames)
		}

		fmt.Println("END###OF###LINE")

	}
	con = db
	return db, nil
}
