package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	// struc do via cep
	// https://viacep.com.br
	// para converter json para struct
	// https://transform.tools/json-to-go
	type ViaCep struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Unidade     string `json:"unidade"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		Uf          string `json:"uf"`
		Estado      string `json:"estado"`
		Regiao      string `json:"regiao"`
		Ibge        string `json:"ibge"`
		Gia         string `json:"gia"`
		Ddd         string `json:"ddd"`
		Siafi       string `json:"siafi"`
	}

	// os.Args é um slice de strings que contém os argumentos da linha de comando
	// os.Args[0] é o nome do programa
	// os.Args[1:] são os argumentos fornecidos pelo usuário
	// os argumentos fornecidos pelo usuário são no momento de executar o programa
	// go run main.go arg1 arg2 arg3

	//fmt.Println(os.Args[0]) // imprime o nome do programa
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/"+ cep + "/json/")
		if err != nil {
			// Fprintf joga a mensagem de erro para algum lugar, nesse caso, para o stderr
			// stderr é o fluxo de saída de erro padrão, que geralmente é o terminal
			fmt.Fprintf(os.Stderr, "Error no request: %v\n", err)
		}

		defer req.Body.Close()

		result, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error na leitura: %v\n", err)
		}

		var data ViaCep

		err = json.Unmarshal(result, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error no unmarshal: %v\n", err)
		}

		fmt.Println(data)

		fmt.Println(string(result))

		file, err := os.Create(cep + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error na criação do arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = file.Write(result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error na escrita do arquivo: %v\n", err)
		}
	}

	time.Sleep(60 * time.Second)


	for _, cep := range os.Args[1:] {

		os.Remove(cep + ".json")
		fmt.Println("Arquivo", cep + ".json", "removido com sucesso!")
	}
}
