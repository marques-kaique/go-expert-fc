package main

import "net/http"

// ServeMux é um multiplexador de solicitações HTTP.
// Ele implementa a interface http.Handler para que possa ser usado como um manipulador em http.ListenAndServe.
// Podendo ser utilizado para criar um roteador simples.
// E ter varias rotas com diferentes manipuladores.
// Assim você possui um controle mais granular sobre como as solicitações são tratadas.
// Caso use o padrão nil, qualquer solicitação para qualquer rota será tratada pelo manipulador padrão http.DefaultServeMux.
// ou seja, se uma biblioteca ou pacote que você está usando chamar http.ListenAndServe(":8080", nil), ele usará o manipulador padrão.
// O que significa que você não terá controle sobre como as solicitações são tratadas.
// injetando endpoints em seu aplicativo sem o seu conhecimento.

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.Handle("/blog", &blog{})

	http.ListenAndServe(":8080", mux)

	// podemos ter n portas
	// e cada uma com seu próprio mux
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":9090", mux2)
}

type blog struct{}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my blog!"))
}
