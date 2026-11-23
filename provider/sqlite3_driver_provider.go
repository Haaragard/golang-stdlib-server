package provider

import (
	"fmt"
	"log"
	"reflect"

	"github.com/haaragard/golang-stdlib-server/container"
	"github.com/haaragard/golang-stdlib-server/infrastructure/database"
)

type Sqlite3DriverProvider struct {
	container *container.HaaraContainer
	driver    *database.DatabaseSqlite3Driver
}

func (provider *Sqlite3DriverProvider) SetContainer(container *container.HaaraContainer) {
	provider.container = container
}

func (provider *Sqlite3DriverProvider) OnStart() {
	fmt.Println("Sqlite3DriverProvider OnStart Starting")
	{
		dsn := "file:database.db?cache=shared"
		provider.driver = database.NewDatabaseSqlite3Driver(dsn)
		if err := provider.driver.Open(); err != nil {
			log.Fatalf("Could not open connection with Sqlite3Driver: %v", err)
		} else {
			provider.container.RegisterAs(reflect.TypeFor[database.DatabaseSql](), provider.driver)
		}
	}
	fmt.Println("Sqlite3DriverProvider OnStart Done")
}

func (provider *Sqlite3DriverProvider) OnShutdown() {
	fmt.Println("Sqlite3DriverProvider OnShutdown Starting")
	provider.driver.Close()
	fmt.Println("Sqlite3DriverProvider OnShutdown Done")
}
