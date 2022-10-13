package rabbitmq

import (
	apiutil "api-rabbit-sender/api_util"
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func rabbitConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	apiutil.FailOnError(err, "Erro de conexao com rabbitmq")
	return conn
}

func EventSender(rabbitConnection *amqp.Connection, body []byte) bool {
	channel, err := rabbitConnection.Channel()
	apiutil.FailOnError(err, "Erro ao gerar channel")
	defer rabbitConnection.Close()
	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		"",
		"golang_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	log.Printf("[*] sent %s", body)
	apiutil.FailOnError(err, "Falha ao enviar msg para fila")

	return true
}
