package main

import (
	"fmt"

	"github.com/marques-kaique/go-expert-fc/aula/packaging/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
