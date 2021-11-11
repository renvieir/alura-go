package routes

import (
	"net/http"

	"alura-go/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
