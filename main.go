package main

import (
	"net/http"

	"alura-go/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
