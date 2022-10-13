package controller

import (
	"api-rabbit-sender/api_util"
	"api-rabbit-sender/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func prepareBody(w http.ResponseWriter, r *http.Request) {
	var person model.Person

	body, err := io.ReadAll(r.Body)
	apiutil.FailOnError(err, "Erro ao converter body da requisição")

	if ConvertBodyToStruct(body, &person) != nil {
		w.Write([]byte("Dados incorretos"))
	} else {

	}

}

func ConvertBodyToStruct(body []byte, person *model.Person) error {
	if err := json.Unmarshal(body, &person); err != nil {
		log.Printf("Erro ao converter body para models")
		return err
	}
	return nil
}
