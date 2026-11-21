package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"

	"github.com/haaragard/stdlib-rest/controller"
	"github.com/haaragard/stdlib-rest/server"
)

func main() {
	hhs := server.NewHaaraHttpServer()
	hhr := hhs.GetRouter()
	hhr.AddRoute("GET", "/health", controller.HealthHandlerFunc)
	hhr.Post("/health", controller.HealthHandlerFunc)
	hhr.Put("/health", controller.HealthHandlerFunc)
	hhr.Patch("/health", controller.HealthHandlerFunc)
	hhr.Delete("/health", controller.HealthHandlerFunc)
	hhr.Get("/health2", controller.HealthHandlerFunc)
	hhr.Get("/health3", controller.HealthHandlerFunc)

	ListAndPrintRegisteredRoutes(hhs)

	log.Fatal(hhs.ServeAndListen())
}

func ListAndPrintRegisteredRoutes(haaraHttpServer *server.HaaraHttpServer) {
	hhr := haaraHttpServer.GetRouter()
	fmt.Printf("\n\nRoute List:\n")
	listRoutes := hhr.GetRoutes()
	for routePath, routeHandlerFunc := range listRoutes {
		pc := reflect.ValueOf(routeHandlerFunc).Pointer()
		fullName := runtime.FuncForPC(pc).Name()

		fmt.Printf("path: \"%s\" | method: \"%s\"\n", routePath, fullName)
	}
	fmt.Printf("\n\n")
}
