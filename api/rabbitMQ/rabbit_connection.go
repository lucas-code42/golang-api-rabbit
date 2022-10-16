package rabbitmq

import (
	apiutil "api-rabbit-sender/api_util"
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// EventSender estabelece uma conexao com rabbit e envia o pacote para fila setada na funcao.
func EventSender(body []byte) bool {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	apiutil.FailOnError(err, "Erro de conexao com rabbitmq")
	defer conn.Close()

	channel, err := conn.Channel()
	apiutil.FailOnError(err, "Erro ao gerar channel")
	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var queue string = "golang_queue"

	err = channel.PublishWithContext(
		ctx,
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	log.Printf("[*] sent %s", body)

	if err != nil {
		return false
	} else {
		return true
	}
}
