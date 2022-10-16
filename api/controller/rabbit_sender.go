package controller

import (
	apiutil "api-rabbit-sender/api_util"
	"api-rabbit-sender/model"
	rabbitMQ "api-rabbit-sender/rabbitMQ"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// PrepareBody recebe a requisão e valida se esta de acordo com o model para entao ser postada msg na fila
func PrepareBody(w http.ResponseWriter, r *http.Request) {
	var person model.Person

	body, err := io.ReadAll(r.Body)
	apiutil.FailOnError(err, "Erro ao converter body da requisição")

	if apiutil.ConvertBodyToStruct(body, &person) != nil {
		w.Write([]byte("Dados incorretos"))
	} else {
		queueData, err := json.Marshal(person)
		if err != nil {
			log.Print("Erro ao fazer Marshal", err)
			return
		}
		rabbitMQ.EventSender(queueData)
	}

}
