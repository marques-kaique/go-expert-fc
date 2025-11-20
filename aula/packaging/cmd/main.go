package main

import (
	"fmt"

	"github.com/ksmarques/go-expert-fc-main/aula/packaging/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
