package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"

	"github.com/haaragard/golang-stdlib-server/application"
	"github.com/haaragard/golang-stdlib-server/controller"
	"github.com/haaragard/golang-stdlib-server/controller/user"
	"github.com/haaragard/golang-stdlib-server/infrastructure/database"
	"github.com/haaragard/golang-stdlib-server/provider"
	"github.com/haaragard/golang-stdlib-server/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Get Application context
	context, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	defer stop()

	hhs := server.NewHaaraHttpServer()
	hhr := hhs.GetRouter()

	RegisterAppRoutes(hhr)
	ListAndPrintRegisteredRoutes(hhr)

	app := application.NewHaaraApplication(hhs)
	app.RegisterProvider(&provider.Sqlite3DriverProvider{})

	app.Run()

	<-context.Done()
	if err := app.Stop(context); err != nil {
		log.Fatalf("could not shutdown: %v\n", err)
	}

	os.Exit(0)
}

func ConnectDatabase() database.DatabaseSql {
	dsn := "file:database.db?cache=shared"
	databaseDriver := database.NewDatabaseSqlite3Driver(dsn)

	err := databaseDriver.Open()
	if err != nil {
		panic(err)
	}
	return databaseDriver
}

func RegisterAppRoutes(haaraHttpRouter *server.HaaraHttpRouter) {
	haaraHttpRouter.Get("/health", controller.HealthHandlerFunc)
	haaraHttpRouter.Post("/users", user.CreateHandlerFunc)
}

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
