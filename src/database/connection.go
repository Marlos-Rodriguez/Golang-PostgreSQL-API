package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//GetConnection ...
func GetConnection() *sql.DB {
	connStr := "postgres://postgres:Elmejordiademivida2009@localhost/todos_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
