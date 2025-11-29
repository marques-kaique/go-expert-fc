package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex{} // Cria um mutex para controlar o acesso à variável number, funciona mas não é a melhor forma de resolver o problema
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock() // Bloqueia o acesso à variável number
		//number++
		//m.Unlock() // Desbloqueia o acesso à variável number, com isso, os valores foram corretamente incrementados
		
		// Incrementa o valor da variável number de forma atômica
		// com essa implementação, não é necessário utilizar um mutex para controlar o acesso à variável number, utilizado para concorrência
		atomic.AddUint64(&number, 1) 
		time.Sleep(300 * time.Microsecond)
		text := fmt.Sprintf("Você teve acesso a essa pagina %d vezes\n", number)
		w.Write([]byte(text))
	})

	http.ListenAndServe(":3000", nil)
}

// Para o teste, foi utilizado o ApacheBench com o seguinte comando: ab -n 10000 -c 100 http://localhost:3000/
// -n -> número de requisições
// 10000 -> número de requisições
// -c -> número de requisições concorrentes
// 100 -> número de requisições concorrentes
// http://localhost:3000/ -> URL do servidor

// Quando executamos novamente o curl localhost:3000/, deveria ser 10000, mas na verdade é um valor inferior
// Ocorrendo nesse caso um problema de concorrência, pois a variável number está sendo acessada por várias goroutines ao mesmo tempo

// go te uma forma de detectar se está ocorrendo uma condição de corrida
// para isso, deve ser utilizado o comando
// go run -race main.go
// caso haja, será exibida uma mensagem de erro indicando a linha do código onde ocorre a condição de corrida
