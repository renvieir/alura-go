package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err.Error())
	}

	return db
}
