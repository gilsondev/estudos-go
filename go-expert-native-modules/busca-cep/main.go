package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const API_URL = "https://viacep.com.br/ws/%s/json/"

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

type CepError struct {
	Erro bool `json:"erro"`
}

func DeserializeCEP(data []byte) (Cep, CepError) {
	var cepError CepError
	var cep Cep
	err := json.Unmarshal(data, &cepError)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao deserializar resposta: %s\n", err)
		os.Exit(1)
	}

	if cepError.Erro {
		return cep, cepError
	}

	err = json.Unmarshal(data, &cep)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao deserializar resposta: %s\n", err)
		os.Exit(1)
	}

	return cep, cepError
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf(API_URL, cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %s\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %s\n", err)
		}

		if req.StatusCode == 400 {
			fmt.Fprintf(os.Stderr, "CEP inválido: %s\n", cep)
		}

		cepData, cepError := DeserializeCEP(res)

		if cepError.Erro {
			fmt.Fprintf(os.Stderr, "CEP não encontrado: %s\n", cep)
			continue
		}

		println("CEP:", cepData.Cep)
		println("Logradouro:", cepData.Logradouro)
		println("Complemento:", cepData.Complemento)
		println("Bairro:", cepData.Bairro)
		println("Localidade:", cepData.Localidade)
		println("UF:", cepData.UF)
		println("GIA:", cepData.GIA)
		println("DDD:", cepData.DDD)
		println("SIAFI:", cepData.SIAFI)
	}
}
