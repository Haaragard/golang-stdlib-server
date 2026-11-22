package database

type DatabaseSql interface {
	Close()
	Get()
}
