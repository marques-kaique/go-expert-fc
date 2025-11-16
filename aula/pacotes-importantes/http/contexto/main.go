package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Contexto é um pacote que permite que passe informações dele para diversa solicitações no sistema
// e pode ser cancelado, por exemplo, quando o usuário cancela a solicitação
// e quando um contexto é cancelado, todas as solicitações que estão utilizando esse contexto também são canceladas
// pode ser trabalhado com contexto utilizando tempo de execução, por exemplo, quando o tempo de execução é muito longo, cancela o contexto
func main() {
	ctx := context.Background() // cria um contexto vazio

	ctx, cancel := context.WithTimeout(ctx, time.Second) // quando alguma coisa for executada utilizando o ctx e demorar mais de 1s, cancela o contexto

	//ctx, cancel := context.WithCancel(ctx) // cria um contexto com cancelamento apenas utilizando o metodo cancel()
	defer cancel() // outra forma de cancelar o contexto, é utilizando o cancel, assim o contexto sempre será cancelado

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://www.google.com", nil) // cria uma nova request com o contexto
	// no lugar de setar o timeout como feito em //timeout, assim o contexto já é criado com o timeout

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req) // utilizando o default client, é o mesmo que utilizar o client http.Client{}

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
