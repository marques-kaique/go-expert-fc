package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	// Estabelece a conexão com o RabbitMQ
	// passando user e senha hardcoded só para exemplo
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// Cria um canal
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queueName string) error {
	msgs, err := ch.Consume(
		queueName,   // queue
		"go-consumer", // consumer
		false,         // auto-ack - somente true em casos específicos, quando pode perder mensagens
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, body string, exchangeName string) error {
	err := ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {

		return err
	}

	return nil
}
