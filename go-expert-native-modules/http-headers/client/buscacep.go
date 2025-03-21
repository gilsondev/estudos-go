package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const API_URL = "https://viacep.com.br/ws/%s/json/"

type ViaCEP struct {
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

type ViaCEPError struct {
	Erro bool `json:"erro"`
}

func BuscaCEP(cep string) (*ViaCEP, error) {
	req, err := http.Get(fmt.Sprintf(API_URL, cep))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	cepData, cepError := deserializeCEP(res)

	if cepError.Erro {
		return nil, fmt.Errorf("CEP %s não encontrado", cep)
	}

	return &cepData, nil
}

func deserializeCEP(data []byte) (ViaCEP, ViaCEPError) {
	var cepError ViaCEPError
	var cep ViaCEP
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
