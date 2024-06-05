package main

import (
	"html/template"
	"net/http"
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
	templates := []string{"header.html", "email.template.html", "footer.html"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("email.template.html").ParseFiles(templates...))
		err := t.Execute(w, Email{
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
	})
	println("Starting server on port 8084")
	http.ListenAndServe(":8084", nil)
}
