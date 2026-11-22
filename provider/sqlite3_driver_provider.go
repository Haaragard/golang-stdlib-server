package provider

import "fmt"

type Sqlite3DriverProvider struct{}

func (driver *Sqlite3DriverProvider) OnStart() {
	fmt.Println("Sqlite3DriverProvider OnStart")
}

func (driver *Sqlite3DriverProvider) OnShutdown() {
	fmt.Println("Sqlite3DriverProvider OnShutdown")
}
