package main

import (
	"fmt"
	"net/http"

	"github.com/gilsondev/aprendagolang-api-pgsql/configs"
	"github.com/gilsondev/aprendagolang-api-pgsql/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handlers.ListHandler)
	mux.HandleFunc("POST /tasks", handlers.CreateHandler)
	mux.HandleFunc("GET /tasks/{id}", handlers.GetHandler)
	mux.HandleFunc("PUT /tasks/{id}", handlers.UpdateHandler)
	mux.HandleFunc("DELETE /tasks/{id}", handlers.DeleteHandler)

	port := fmt.Sprintf(":%s", configs.GetServerPort())
	http.ListenAndServe(port, mux)
}
