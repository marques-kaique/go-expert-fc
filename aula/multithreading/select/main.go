package main

import (
	"time"
)

type Message struct {
	Id  int
	Txt string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	go func() {
		time.Sleep(2 * time.Second)
		msg := Message{Id: 1, Txt: "Hello from RabbitMQ"}
		c1 <- msg
	}()

	go func() {
		time.Sleep(1 * time.Second)
		msg := Message{Id: 2, Txt: "Hello from Kafka"}
		c2 <- msg
	}()

	// nesse cenario, imagina que há duas chamadas para API diferentes
	// o select aguarda a resposta e retorna a aquela que chegar primeiro
	// é possivel por o select dentro de um loop para ficar aguardando as respostas for {} - loop infinito
	for {
		select {
		case msg := <-c1:
			println("Received", msg.Txt)

		case msg := <-c2:
			println("Received", msg.Txt)

		case <-time.After(3 * time.Second): // se nenhuma das outras cases for atendida em 3 segundos, cai aqui
			println("Timeout: no messages received within 3 seconds")

			//default: // default passa pela condição e caso nenhuma seja atendida, cai aqui
			//	println("No messages received")
		}
	}
}
