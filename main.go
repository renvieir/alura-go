package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func connectDB() *sql.DB {
	conn := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable search_path=alura_store"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()

	allProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	products := []Produto{}

	for allProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := allProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		products = append(products, p)
	}
	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}
