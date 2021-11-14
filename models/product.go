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

	allProducts, err := conn.Query("SELECT * FROM products ORDER BY id")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	products := []Produto{}

	for allProducts.Next() {
		err := allProducts.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			panic(err.Error())
		}
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

func DeleteProduct(id string) {
	conn := db.CreateConnection()
	deleteQuery, err := conn.Prepare("DELETE FROM alura_store.products WHERE id=$1;")
	if err != nil {
		panic(err.Error())
	}
	deleteQuery.Exec(id)
	defer conn.Close()
}

func UpdateProduct(id string, nome, descricao string, preco float64, quantidade int) {
	conn := db.CreateConnection()
	updateQuery, err := conn.Prepare("UPDATE alura_store.products SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5;")
	if err != nil {
		panic(err.Error())
	}
	updateQuery.Exec(nome, descricao, preco, quantidade, id)
	defer conn.Close()
}

func FindProductById(id string) Produto {
	conn := db.CreateConnection()
	findQuery, err := conn.Query("SELECT * FROM alura_store.products WHERE id=$1;", id)
	if err != nil {
		panic(err.Error())
	}

	product := Produto{}

	for findQuery.Next() {
		err = findQuery.Scan(&product.Id, &product.Nome, &product.Descricao, &product.Preco, &product.Quantidade)
		if err != nil {
			panic(err.Error())
		}
	}

	defer conn.Close()
	return product
}
