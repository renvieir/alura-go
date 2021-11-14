package main

import (
	"net/http"
	"os"

	"alura-go/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
