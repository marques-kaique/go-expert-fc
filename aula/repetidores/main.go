package main

import "fmt"

func main() {

	//array
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	println("\nLoop for incremental com array")
	for i := 0; i < len(array); i++ {
		println(array[i])
	}

	println("\nLoop for range (foreach)")
	for i, v := range array {
		fmt.Printf("posição: %d valor: %d\n", i, v)
	}

	temp := 0
	x := 0

	println("\nLoop while")
	for {
		x++

		if temp == 0 || temp%2 == 0 {
			fmt.Printf("%d é igual a par\n", temp)
			temp++
		}

		fmt.Printf("%d é igual a impar\n", temp)
		temp++

		if temp == 10 {
			fmt.Printf("Loop interrompido, total de %d iterações\n", x)
			temp = 0
			x = 0
			break
		}
	}

	println("\nLoop while with continue")
	for {
		x++

		if temp == 0 || temp%2 == 0 {
			fmt.Printf("%d é igual a par\n", temp)
			temp++
			continue
		}

		fmt.Printf("%d é igual a impar\n", temp)
		temp++

		if temp == 10 {
			fmt.Printf("Loop interrompido, total de %d iterações\n\n", x)
			break
		}
	}

	//slice
	println("\nLoop for range (foreach) com slice")
	slice := []int{1, 2}
	for i, v := range slice {
		fmt.Printf("posição: %d valor: %d capacidade: %d\n", i, v, cap(slice))

		if cap(slice) == 2 {
			slice = append(slice, 3, 4)
			// mesmo aumentando a capacidade do slice, o loop for range não é afetado.
			// o loop for range é baseado no tamanho do slice no momento em que é iniciado.
			// se o slice for alterado durante o loop, os novos elementos não serão percorridos.
			// para percorrer os novos elementos, é necessário iniciar um novo loop for range.
		}
	}

	//map
	println("\nLoop for range (foreach) com map")
	mapa := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	for k, v := range mapa {
		fmt.Printf("chave: %s valor: %d\n", k, v)
	}

	//string
	println("\nLoop for range (foreach) com string")
	texto := "Olá, mundo!"
	for i, v := range texto {
		fmt.Printf("posição: %d valor: %c\n", i, v)
	}
}
