package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome  string
	Preco float64
}

type Cursos []Curso

func main() {
	curso := Curso{
		Nome:  "Golang",
		Preco: 49.90,
	}

	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmp.Execute(os.Stdout, Cursos{
		curso,
		{Nome: "Python", Preco: 39.90},
		{Nome: "Java", Preco: 29.90},
	})
	if err != nil {
		panic(err)
	}
}
