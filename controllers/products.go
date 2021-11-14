package controllers

import (
	"alura-go/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
		return
	}

	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	precoFloat, err := strconv.ParseFloat(preco, 64)

	if err != nil {
		log.Println("Erro na conversao de preco:", err)
	}

	quantidadeInt, err := strconv.Atoi(quantidade)

	if err != nil {
		log.Println("Erro na conversao de quantidade:", err)
	}

	models.InsertProduct(nome, descricao, precoFloat, quantidadeInt)

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindProductById(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
	id := r.FormValue("id")
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")
	precoFloat, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Erro na conversao de preco:", err)
	}
	quantidadeInt, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Println("Erro na conversao de quantidade:", err)
	}

	models.UpdateProduct(id, nome, descricao, precoFloat, quantidadeInt)
	http.Redirect(w, r, "/", 301)
}
