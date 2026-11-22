package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"

	"github.com/haaragard/golang-stdlib-server/container"
	"github.com/haaragard/golang-stdlib-server/controller"
	"github.com/haaragard/golang-stdlib-server/controller/user"
	"github.com/haaragard/golang-stdlib-server/infrastructure/database"
	"github.com/haaragard/golang-stdlib-server/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	container := container.NewHaaraContainer()

	databasePool := ConnectDatabase()
	defer databasePool.Close()

	container.RegisterDynamic(databasePool)

	hhs := server.NewHaaraHttpServer()
	hhr := hhs.GetRouter()

	RegisterAppRoutes(hhr)
	ListAndPrintRegisteredRoutes(hhr)

	log.Fatal(hhs.ServeAndListen())
}

func ConnectDatabase() database.DatabaseSql {
	dsn := "file:database.db?cache=shared"
	pool, err := database.NewDatabaseSqlite3Driver(dsn)
	if err != nil {
		panic(err)
	}

	return pool
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
