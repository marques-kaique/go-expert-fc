package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5) // valor abaixo de 5s, cancela o contexto
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// select funcionando de forma assincrona
	// fica aguardando o resultado e quando ele chega, executa o case correspondente
	// talvez seja igual observer pattern
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(time.Second * 5):
		fmt.Println("Hotel booking successful.")
	}
}
