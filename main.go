package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"

	"github.com/haaragard/golang-stdlib-server/container"
	"github.com/haaragard/golang-stdlib-server/controller"
	"github.com/haaragard/golang-stdlib-server/controller/user"
	"github.com/haaragard/golang-stdlib-server/server"
)

func main() {
	hhs := server.NewHaaraHttpServer()
	hhr := hhs.GetRouter()

	RegisterAppRoutes(hhr)
	ListAndPrintRegisteredRoutes(hhs)

	TestHaaraContainerInstances()

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

type TesTruct struct {
	ValueOne string
	ValueTwo int
}

func TestHaaraContainerInstances() {
	container := container.NewHaaraContainer()

	// Registers testTruct1 in container
	container.RegisterDynamic(&TesTruct{
		ValueOne: "teste 1",
		ValueTwo: 1,
	})

	// Registers tesTruct2 in container, overriding tesTruct1
	container.RegisterDynamic(&TesTruct{
		ValueOne: "teste 2",
		ValueTwo: 2,
	})

	// The "*TesTruct" (*) is necessary, because we registered a pointer (&TesTruct).
	// The ".(*TesTruct)" is necessary too, because we are causing a conversion from "any"
	// to "*TesTruct" that was registered.
	tesTructReturned, ok := container.Resolve(reflect.TypeFor[*TesTruct]()).(*TesTruct)
	if !ok {
		fmt.Printf("Values does not match")
		return
	}

	fmt.Printf("\n\n\nContainer registration and retrieval test:\n")
	fmt.Printf(
		"the tesTruct returned has the values {valueOne: %s} and {valueTwo: %d}\n",
		tesTructReturned.ValueOne,
		tesTructReturned.ValueTwo,
	)
}
