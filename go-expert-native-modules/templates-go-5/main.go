package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
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

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	http.HandleFunc("/", CursosHandler)
	http.ListenAndServe(":8080", nil)
}

func CursosHandler(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(w, Cursos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao renderizar template: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
