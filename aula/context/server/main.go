package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request initiated")
	w.Write([]byte("Request initiated\n"))

	defer log.Println("Request completed")

	select {
	case <-time.After(time.Second * 5):
		// imprime no terminal
		log.Println("Request  processed with success")
		// imprime no browser
		w.Write([]byte("Request processed with success\n"))
	case <-ctx.Done(): // quando o contexto é cancelado, o case ctx.Done() é executado
		log.Println("Request cancelled in the client")
	}
}
