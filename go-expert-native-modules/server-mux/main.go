package main

import (
	"fmt"
	"net/http"
	"servermux/routes"
)

const API_URL = "https://viacep.com.br/ws/%s/json/"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.CEPHandler)

	fmt.Println("Iniciando servidor na porta 8080")
	http.ListenAndServe(":8080", mux)
}
