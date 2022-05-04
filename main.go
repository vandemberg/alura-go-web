package main

import (
	"alura-go-web/routes"
	"net/http"
)

func main() {
	routes.CarregaRoutas()
	http.ListenAndServe(":8000", nil)
}
