package main

import (
	"api-rabbit-sender/routes"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	msg := figure.NewColorFigure("api  rabbitMQ  sender", "", "green", true)
	msg.Print()

	routes.RoutesHandler()
}
