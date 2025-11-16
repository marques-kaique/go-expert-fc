package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}

	// body é um io.Reader, sendo um slice de bytes
	// mesmo que []byte seja um slice de bytes, é necessário utilizar o bytes.NewBuffer que devolve um buffer de bytes
	// buffer de bytes é um io.Reader
	jsonVar := bytes.NewBuffer([]byte(`{"nome": "Curso de Go", "preco": 29.90}`))

	resp, err := client.Post("http://www.google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	// pega os dados do body e escreve no stdout
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
