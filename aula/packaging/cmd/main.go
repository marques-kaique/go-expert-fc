package main

import (
	"fmt"

	"github.com/marques-kaique/go-expert-fc/aula/packaging/math" //recomendado usar sempre url completa
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())

	fmt.Println(math.Generate())
}
