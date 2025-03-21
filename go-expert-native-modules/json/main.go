package main

import (
	"encoding/json"
	"os"
	"strings"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func CreateJSONViaMarshal() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println("Creating JSON via Marshal")
	println(string(res))
}

func CreateJSONViaEncoder() {
	conta := Conta{Numero: 1, Saldo: 100}

	println("Creating JSON via Encoder")
	json.NewEncoder(os.Stdout).Encode(conta)
}

func ReadJSONViaUnmarshal() {
	data := `{"numero": 2, "saldo": 200}`
	var conta Conta
	err := json.Unmarshal([]byte(data), &conta)
	if err != nil {
		panic(err)
	}

	println("Reading JSON via Unmarshal")
	println("Número:", conta.Numero)
	println("Saldo:", conta.Saldo)
}

func ReadJSONViaDecoder() {
	data := `{"numero": 3, "saldo": 300}`
	var conta Conta
	decoder := json.NewDecoder(strings.NewReader(data))
	err := decoder.Decode(&conta)
	if err != nil {
		panic(err)
	}

	println("Reading JSON via Decoder")
	println("Número:", conta.Numero)
	println("Saldo:", conta.Saldo)
}

func main() {
	CreateJSONViaMarshal()
	CreateJSONViaEncoder()

	ReadJSONViaUnmarshal()
	ReadJSONViaDecoder()
}
