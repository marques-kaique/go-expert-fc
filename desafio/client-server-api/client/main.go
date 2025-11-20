package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type Bid struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data Bid
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	f, err := os.Create("./cotacaoe.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	text := "Dólar: " + data.Bid
	_, err = f.Write([]byte(text))
	if err != nil {
		panic(err)
	}
}
