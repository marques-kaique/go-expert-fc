package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	ID    int     `json:"a"` // isso é uma tag, é um metadado que pode ser usado para serializar e deserializar
	Saldo float64 `json:"b"` // "-" ignora o campo
}

func main() {
	conta := Conta{
		ID:    1,
		Saldo: 100.00,
	}

	// marshal converte um valor para JSON
	// sendo um serilizador
	// serializador é um pacote que converte um valor para outro formato
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	// resultado é um slice de bytes
	fmt.Println(string(res))

	// encode para um arquivo ou tela do terminal

	enconder := json.NewEncoder(os.Stdout) // cria um enconder
	err = enconder.Encode(conta)           // converte a conta para JSON e escreve no terminal
	if err != nil {
		panic(err)
	}

	// poderia ser feito assim também
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	var contaX Conta

	// unmarshal converte um JSON para um valor
	// sendo um deserializador
	// deserializador é um pacote que converte um valor para outro formato

	err = json.Unmarshal(res, &contaX) // converte o res para contaX
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID : %d, Saldo: %v\n", contaX.ID, contaX.Saldo)

	jsonString := []byte(`{"a": 2, "b": 0.00}`)
	var contaY Conta
	err = json.Unmarshal(jsonString, &contaY)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID : %d, Saldo: %v\n", contaY.ID, contaY.Saldo)
}
