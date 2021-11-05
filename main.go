package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func connectDB() *sql.DB {
	conn := "user=postgres dbname=postgrespassword=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tênis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}

func main() {
	db := connectDB()
	defer db.Close()
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}
