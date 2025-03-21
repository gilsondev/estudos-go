package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{Nome: "Golang", CargaHoraria: 40}
	content := "Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"
	t := template.Must(template.New("CursoTemplate").Parse(content))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
