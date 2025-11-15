package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome  string
	Preco float64
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", CursoTemplateHandler)

	http.ListenAndServe(":8080", nil)
}

func CursoTemplateHandler(w http.ResponseWriter, r *http.Request) {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	curso := Curso{
		Nome:  "Golang",
		Preco: 49.90,
	}

	tmp := template.Must(template.New("content.html").ParseFiles(templates...))

	err := tmp.Execute(w, Cursos{
		curso,
		{Nome: "Python", Preco: 39.90},
		{Nome: "Java", Preco: 29.90},
		{Nome: "C#", Preco: 19.90},
		{Nome: "Ruby", Preco: 59.90},
		{Nome: "JavaScript", Preco: 4.90},
	})
	if err != nil {
		panic(err)
	}
}
