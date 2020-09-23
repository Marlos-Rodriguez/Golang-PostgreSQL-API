package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//GetConnection ...
func GetConnection() *sql.DB {
	URLDB := os.Getenv("DATABASE_URL")

	connStr := URLDB
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
