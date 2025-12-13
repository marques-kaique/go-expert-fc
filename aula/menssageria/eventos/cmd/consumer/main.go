package main

import (
	"fmt"

	"github.com/marques-kaique/go-expert-fc/aula/menssageria/eventos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		// Processa a mensagem recebida
		fmt.Printf("Received a message: %s\n", msg.Body)
		// Confirma o recebimento da mensagem
		// O false indica que não precisa por novamente a mensagem na fila
		msg.Ack(false)
	}
}
