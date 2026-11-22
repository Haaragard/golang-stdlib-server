package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/haaragard/golang-stdlib-server/application"
	"github.com/haaragard/golang-stdlib-server/controller"
	"github.com/haaragard/golang-stdlib-server/controller/user"
	"github.com/haaragard/golang-stdlib-server/provider"
	"github.com/haaragard/golang-stdlib-server/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	hhs := server.NewHaaraHttpServer()
	hhr := hhs.GetRouter()

	RegisterAppRoutes(hhr)
	ListAndPrintRegisteredRoutes(hhr)

	app := application.NewHaaraApplication(hhs)
	app.RegisterProvider(&provider.Sqlite3DriverProvider{})

	// Runs application and block main runtime thread.
	app.Run()
}

func RegisterAppRoutes(haaraHttpRouter *server.HaaraHttpRouter) {
	haaraHttpRouter.Get("/health", controller.HealthHandlerFunc)
	haaraHttpRouter.Post("/users", user.CreateHandlerFunc)
}

// Helper to enable registered routes visualization
func ListAndPrintRegisteredRoutes(haaraHttpRouter *server.HaaraHttpRouter) {
	fmt.Printf("\n\nRoute List:\n")
	listRoutes := haaraHttpRouter.GetRoutes()
	for routePath, routeHandlerFunc := range listRoutes {
		pc := reflect.ValueOf(routeHandlerFunc).Pointer()
		fullName := runtime.FuncForPC(pc).Name()

		fmt.Printf("path: \"%s\" | method: \"%s\"\n", routePath, fullName)
	}
	fmt.Printf("\n\n")
}
