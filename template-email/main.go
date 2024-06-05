package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Name        string
	Description string
}

type Email struct {
	Name    string
	Courses []Curso
}

func main() {
	t := template.Must(template.New("email.template.html").ParseFiles("email.template.html"))
	err := t.Execute(os.Stdout, Email{
		Name: "John Doe",
		Courses: []Curso{
			{"Golang", "Learn Golang"},
			{"Python", "Learn Python"},
			{"Java", "Learn Java"},
		},
	})
	if err != nil {
		panic(err)
	}
}
