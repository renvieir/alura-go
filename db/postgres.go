package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	conection := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable search_path=alura_store"
	db, err := sql.Open("postgres", conection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
