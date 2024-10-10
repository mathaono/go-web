package main

import (
	"go-web/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	//Func importada do package routes/routes.go
	routes.LoadingRoutes()
	http.ListenAndServe(":8000", nil)
}
