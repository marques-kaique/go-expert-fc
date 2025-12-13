package main

import "github.com/marques-kaique/go-expert-fc/aula/menssageria/eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, RabbitMQ!", "amq.direct")
}
