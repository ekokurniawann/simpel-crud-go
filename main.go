package main

import (
	"crud-simpel/routes"
	"net/http"
)

func main() {
	router := routes.RegisterRoutes()

	http.ListenAndServe(":8080", router)
}
