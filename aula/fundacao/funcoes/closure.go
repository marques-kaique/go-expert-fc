package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func closure() {
	// função anônima
	// função que retorna uma função
	// função que roda dentro de uma função
	total := func() int {
		return sum(1, 2)
	}()

	fmt.Println(total)
}
