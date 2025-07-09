package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome  string
	Preco float64
}

func main() {
	curso := Curso{
		Nome:  "Golang",
		Preco: 49.90,
	}

	tmp := template.New("CursoTemoplate")

	// para incluir o valor do campo Nome e Preco no template, use {{.Nome}} e {{.Preco}} (.)
	tmp, _ = tmp.Parse("Curso: {{.Nome}} Preço: {{.Preco}}")

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
