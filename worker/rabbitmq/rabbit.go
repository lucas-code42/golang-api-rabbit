package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"rabbit-worker/controller"
	"rabbit-worker/db"
	model "rabbit-worker/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Worker responsavel por estabelecer conexao com rabbit e persisir dados recebidos no banco de dados
func Worker() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Panic("Erro de conexao com rabbit", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic("Erro ao estabelecer canal", err)
	}
	defer ch.Close()

	var queue string = "golang_queue"

	msg, err := ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic("Erro ao consumir dados", err)
	}

	var dataModel chan model.DataRabbit
	var dataTest model.DataRabbit

	connect, err := db.ConnectDB()
	if err != nil {
		log.Panic("erro de conexao com banco")
	}
	defer connect.Close()

	go func() {
		for msgPack := range msg {

			data, err := ConvertBodyToStruct(msgPack.Body, &dataTest)
			if err != nil {
				fmt.Println("erro ao converter msg da fila para struct")
			}
			fmt.Println("msg recebida", data)

			perssit := controller.PersistData(connect, data)
			if !perssit {
				log.Panic("Erro ao persistir dados no banco")
			} else {
				log.Print("Sucesso!")
			}

		}
	}()

	fmt.Println("[*] Aguardando msgs")
	<-dataModel

}

// ConvertBodyToStruct converte uma requisicao web para um struct desejado, nesse caso esta setado para o unico model da api
func ConvertBodyToStruct(body []byte, person *model.DataRabbit) (model.DataRabbit, error) {
	if err := json.Unmarshal(body, &person); err != nil {
		log.Printf("Erro ao converter body para models")
		return model.DataRabbit{}, err
	}
	return *person, nil
}
