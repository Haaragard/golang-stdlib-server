package database

type DatabaseSql interface {
	Open() error
	Close()
	IsOpen() bool
	Get()
}
