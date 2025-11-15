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

	// template.Must() é uma função que retorna um template e um erro. Se o erro for diferente de nil, ele irá chamar panic(err)
	// must é usado para garantir que o template seja compilado corretamente e que não haja erros de sintaxe.
	// Se houver um erro, o programa irá parar e exibir o erro.
	tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} Preço: {{.Preco}}"))

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
