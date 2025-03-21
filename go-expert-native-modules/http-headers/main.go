package main

import (
	"buscacepv1/client"
	"encoding/json"
	"fmt"
	"net/http"
)

const API_URL = "https://viacep.com.br/ws/%s/json/"

func main() {
	fmt.Println("Iniciando servidor na porta 8080")
	http.HandleFunc("/", CEPHandler)
	http.ListenAndServe(":8080", nil)
}

func CEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Não encontrado"))
		return
	}

	cep := r.URL.Query().Get("cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP não informado"))
		return
	}

	cepData, err := client.BuscaCEP(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Erro ao buscar CEP: %s", err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cepData)
}
