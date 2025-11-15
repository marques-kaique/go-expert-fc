package main

import (
	"fmt"
)

func forceErrorArray() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de um erro:", r)
			// Este erro ocorre porque tentamos acessar o índice 3 de um array com tamanho 3.
			// Índices válidos para [3]int são: 0, 1, 2.
			// Arrays são de tamanho fixo, então não é possível adicionar um novo elemento a um array existente.
		}
	}()

	var array [3]int

	println("*** Array ***")

	fmt.Printf("\ncapacidade: %d, %v tamanho: %v\n\n", cap(array), array, len(array))

	temp := 0

	for {
		array[temp] = temp
		temp++
	}
}

func workWithSlice() {
	slice := []int{1, 2}

	println("*** Slice inicial ***")

	fmt.Printf("\ncapacidade: %d, %v tamanho: %v\n\n", cap(slice), slice, len(slice))
	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice[:0]), slice[:0], len(slice[:0]))

	slice = append(slice, 3, 4)

	println("*** Duplicando valores do Slice ***")

	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice), slice, len(slice))
	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice[:2]), slice[:2], len(slice[:2]))

	slice = append(slice, 5)
	
	println("*** Agregando apenas um novo valor ***")

	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice), slice, len(slice))
	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice[2:]), slice[2:], len(slice[2:]))

	slice = append(slice, 5, 6, 7, 8)

	println("*** Agregando valores além da capacidade anterior ***")

	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice), slice, len(slice))
	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice[2:4]), slice[2:4], len(slice[2:4]))

	// quanto mais elementos adicionamos ao slice, maior é a capacidade do slice.
	// slice possui crescimento exponencial.

	sliceTemp := make([]int, 2)

	println("*** Slice temporário inicializado sem valores ***")
	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(sliceTemp), sliceTemp, len(sliceTemp))

	sliceTemp = append(sliceTemp, 1, 2, 3, 4, 5)

	println("*** Agregando valores ao slice temporário ***")
	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(sliceTemp), sliceTemp, len(sliceTemp))
	
	println("O make alocou elementos em memoria para o sliceTemp")
	println("Ao agregar novos valores ao sliceTemp, os valores anteriores não são perdidos")

	println("\n*** Agregando os slices ***")

	slice = append(slice, sliceTemp...)
	fmt.Printf("capacidade: %d, %v tamanho: %v\n\n", cap(slice), slice, len(slice))
}

func main() {
	forceErrorArray()
	workWithSlice()
}
