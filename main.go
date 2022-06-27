package main

import (
	"net/http"

	"alura.com/app_web/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
