package apiutil

import (
	"api-rabbit-sender/model"
	"encoding/json"
	"log"
)

// FailOnError funcao que recebe um erro e uma msg, caso erro realmente exista o programa entra em panico
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%v ---> %s", err, msg)
	}
}

// ConvertBodyToStruct converte uma requisicao web para um struct desejado, nesse caso esta setado para o unico model da api
func ConvertBodyToStruct(body []byte, person *model.Person) error {
	if err := json.Unmarshal(body, &person); err != nil {
		log.Printf("Erro ao converter body para models")
		return err
	}
	return nil
}
