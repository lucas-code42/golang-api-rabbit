package routes

import (
	"log"
	"net/http"

	"api-rabbit-sender/controller"

	"github.com/gorilla/mux"
)

func RoutesHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Hello).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":5000", r))
}
