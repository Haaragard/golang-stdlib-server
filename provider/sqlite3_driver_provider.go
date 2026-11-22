package provider

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/haaragard/golang-stdlib-server/container"
	"github.com/haaragard/golang-stdlib-server/infrastructure/database"
)

type Sqlite3DriverProvider struct {
	container *container.HaaraContainer
	driver    *database.DatabaseSqlite3Driver
	wg        *sync.WaitGroup
}

func (provider *Sqlite3DriverProvider) SetWaitGroup(wg *sync.WaitGroup) {
	provider.wg = wg
}

func (provider *Sqlite3DriverProvider) SetContainer(container *container.HaaraContainer) {
	provider.container = container
}

func (provider *Sqlite3DriverProvider) OnStart() {
	defer provider.done()

	fmt.Printf("Sqlite3DriverProvider OnStart Starting\n")
	{
		dsn := "file:database.db?cache=shared"
		provider.driver = database.NewDatabaseSqlite3Driver(dsn)
		if err := provider.driver.Open(); err != nil {
			log.Fatalf("Could not open connection with Sqlite3Driver: %v", err)
		} else {
			provider.container.RegisterAs(reflect.TypeFor[database.DatabaseSql](), provider.driver)
		}
	}
	fmt.Printf("Sqlite3DriverProvider OnStart Done\n\n")
}

func (provider *Sqlite3DriverProvider) OnShutdown() {
	defer provider.done()

	fmt.Printf("Sqlite3DriverProvider OnShutdown Starting\n")
	provider.driver.Close()
	fmt.Printf("Sqlite3DriverProvider OnShutdown Done\n\n")
}

func (provider *Sqlite3DriverProvider) done() {
	provider.wg.Done()
}
