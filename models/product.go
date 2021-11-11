package models

import "alura-go/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SearchAllProducts() []Produto {
	conn := db.CreateConnection()

	allProducts, err := conn.Query("select * from products")

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
	defer conn.Close()

	return products
}

func InsertProduct(nome, descricao string, preco float64, quantidade int) {
	conn := db.CreateConnection()
	insertQuery, err := conn.Prepare("INSERT INTO alura_store.products (nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4);")
	if err != nil {
		panic(err.Error())
	}
	insertQuery.Exec(nome, descricao, preco, quantidade)
	defer conn.Close()
}
