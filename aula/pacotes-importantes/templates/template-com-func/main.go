package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome  string
	Preco float64
}

type Cursos []Curso

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func main() {

	http.HandleFunc("/", CursoTemplateHandler)

	http.ListenAndServe(":8080", nil)
}

func CursoTemplateHandler(w http.ResponseWriter, r *http.Request) {

	templates := []string{
		"../compondo-templates/header.html",
		"content.html",
		"../compondo-templates/footer.html",
	}

	curso := Curso{
		Nome:  "Golang",
		Preco: 49.90,
	}

	// utilizando funcs, é necessário criar um novo template e adicionar as funções a ele
	// após isso, é necessário fazer o parse dos arquivos de template
	// não é possível utilizar o template.Must() com o template.New()
	tmp := template.New("content.html")
	tmp.Funcs(template.FuncMap{
		"ToUpper": Uppercase,
	})
	tmp = template.Must(tmp.ParseFiles(templates...))

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
