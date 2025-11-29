package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/martian/v3/log"
)

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCep struct {
	Cep        string `json:"cep"`
	Uf         string `json:"uf"`
	Localidade string `json:"localidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	brasilApi := make(chan BrasilApi)
	viaCep := make(chan ViaCep)

	go func() {
		getCepBrasilApi(brasilApi, ctx)
	}()

	go func() {
		getCepViaCep(viaCep, ctx)
	}()

	for {
		select {
		case res := <-brasilApi:
			fmt.Println("BrasilApi: ", res)
			cancel()
			return

		case res := <-viaCep:
			fmt.Println("ViaCep: ", res)
			cancel()
			return

		case <-time.After(time.Second):
			fmt.Println("Timeout: não foi possível obter os dados a tempo")
			cancel()
			return
		}
	}
}

func getCepBrasilApi(cep chan<- BrasilApi, ctx context.Context) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://brasilapi.com.br/api/cep/v1/01153000", nil)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Errorf("Error no request: %v\n", err)
		return
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error na leitura: %v\n", err)
		return
	}

	defer resp.Body.Close()

	var data BrasilApi

	err = json.Unmarshal(result, &data)

	if err != nil {
		log.Errorf("Error no unmarshal: %v\n", err)
		return
	}

	cep <- data
}

func getCepViaCep(cep chan<- ViaCep, ctx context.Context) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/01153000/json", nil)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Errorf("Error no request: %v\n", err)
		return
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error na leitura: %v\n", err)
		return
	}

	defer resp.Body.Close()

	var data ViaCep

	err = json.Unmarshal(result, &data)

	if err != nil {
		log.Errorf("Error no unmarshal: %v\n", err)
		return
	}

	cep <- data
}
