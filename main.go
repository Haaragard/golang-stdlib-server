package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"

	"github.com/haaragard/stdlib-rest/controller"
	"github.com/haaragard/stdlib-rest/controller/user"
	"github.com/haaragard/stdlib-rest/server"
)

func main() {
	hhs := server.NewHaaraHttpServer()
	hhr := hhs.GetRouter()

	RegisterAppRoutes(hhr)
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

func RegisterAppRoutes(haaraHttpRouter *server.HaaraHttpRouter) {
	haaraHttpRouter.Get("/health", controller.HealthHandlerFunc)
	haaraHttpRouter.Post("/users", user.CreateHandlerFunc)
}
