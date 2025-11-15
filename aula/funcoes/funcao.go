package main

import (
	"errors"
	"fmt"
)

func soma(a, b int) int { // a e b são parâmetros - int é o tipo de retorno
	return a + b
}

func divisao(a, b int) (int, int) { //retorna dois valores
	return a / b, a % b
}

func divisaoComZero(a, b int) (int, error) { // padrão de retorno de erro ser o último
	if b == 0 {
		return 0, errors.New("ERRO: Não é possível dividir por zero")
	}

	return a / b, nil
}

func somandoNValores(valores ...int) int {
	result := 0
	
	for _, v := range valores {
		result += v
	}

	return result
}

func concatenandoComNValores(nome string, valores ...int) string {
	
	for _, v := range valores {
		// fmt.Sprint converte o valor para string
		// existe o strconv.Itoa que faz a mesma coisa
		// https://pkg.go.dev/strconv@go1.24.3
		nome = nome + fmt.Sprint(v)
	}

	return nome
}

func funcao() {
	println("*** Soma")
	println(soma(1, 2))
	
	println("\n*** Divisão com resto")
	println(divisao(10, 3))

	println("\n*** Divisão com zero e geração de erro")
	result, err := divisaoComZero(10, 0)
	if err != nil {
		// go não tem exceções, mas tem erros
		// erros sempre devem ser tratados
		// não possui try-catch
		println(err.Error())
		// para parar a execução do programa
		// os.Exit(1)
		// ou
		//return - utilizado em funções
	}
	println(result)

	println("\n*** Somando N valores")
	println(somandoNValores(1, 2, 3, 4, 5))

	println("\n*** Concatenando com N valores")
	println(concatenandoComNValores("teste", 1, 2, 3, 4, 5, 4))
}
