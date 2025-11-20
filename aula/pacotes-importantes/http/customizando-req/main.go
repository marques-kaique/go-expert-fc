package main

import (
	"io"
	"net/http"
)

func main() {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	// informa que quando receber a resposta, aceita o formato json
	req.Header.Set("Accept", "application/json")

	// meu client, vai receber a request no metodo Do, enviando a request e retornado o response e o erro
	// isso é uma forma de customizar a request, utilizado quando quer passar headers, query params, etc
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
