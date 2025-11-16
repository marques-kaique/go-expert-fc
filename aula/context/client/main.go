package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

// para rodar, executar o server primeiro
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // inferior a 5s cancela o contexto
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// copiar o body da resposta para o stdout -> terminal
	io.Copy(os.Stdout, res.Body)
}
