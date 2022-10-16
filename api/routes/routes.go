package routes

import (
	"log"
	"net/http"

	"api-rabbit-sender/controller"

	"github.com/gorilla/mux"
)

// RoutesHandler controla o fluxo e direcionamento de rotas da api
func RoutesHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Hello).Methods(http.MethodGet)
	r.HandleFunc("/rabbit", controller.PrepareBody).Methods((http.MethodGet))

	log.Fatal(http.ListenAndServe(":5000", r))
}
