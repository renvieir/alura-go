package main

import (
	"net/http"
	"text/template"

	"alura-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
