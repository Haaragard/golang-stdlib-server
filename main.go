package main

import (
	"log"

	"github.com/haaragard/stdlib-rest/controller"
	"github.com/haaragard/stdlib-rest/server"
)

func main() {
	hhs := server.NewHaaraHttpServer()
	hhs.AddRoute("GET", "/health", controller.HealthHandlerFunc)
	log.Fatal(hhs.ServeAndListen())
}
