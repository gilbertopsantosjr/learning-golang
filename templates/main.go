package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	println("Starting server on port 8083")
	curso := Curso{"Golang", 40}
	// t := template.Must(template.New("CursoTemplate").Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas."))
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas.")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
