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
	curso := Curso{
		Nome:  "Golang",
		Preco: 49.90,
	}

	tmp := template.Must(template.New("template.html").ParseFiles("../arquivo-externo/template.html"))

	err := tmp.Execute(w, Cursos{
		curso,
		{Nome: "Python", Preco: 39.90},
		{Nome: "Java", Preco: 29.90},
	})
	if err != nil {
		panic(err)
	}
}
