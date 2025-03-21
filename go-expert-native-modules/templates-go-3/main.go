package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

var Cursos = []Curso{
	{Nome: "Golang", CargaHoraria: 40},
	{Nome: "Java", CargaHoraria: 30},
	{Nome: "C#", CargaHoraria: 20},
	{Nome: "C++", CargaHoraria: 10},
	{Nome: "Javascript", CargaHoraria: 50},
	{Nome: "Typescript", CargaHoraria: 60},
	{Nome: "Python", CargaHoraria: 70},
}

func main() {
	http.HandleFunc("/", CursosHandler)
	http.ListenAndServe(":8080", nil)
}

func CursosHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template.html"))
	err := t.Execute(w, Cursos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao renderizar template: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
