package controllers

import (
	"alura-go/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
