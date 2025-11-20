package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "uuid")
	bookHotel(ctx, "any")
}

func bookHotel(ctx context.Context, name string) { // por comvenção, o primeiro parametro sempre é o contexto
	token := ctx.Value("token")
	fmt.Println(token, " ", name)
}
