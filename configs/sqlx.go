package configs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// get reader connection
func GetReaderSqlx() *sqlx.DB {
	reader := sqlx.MustConnect("mysql", DB_CONNECTION)
	return reader
}

// get writter connection
func GetWriterSqlx() *sqlx.DB {
	writer := sqlx.MustConnect("mysql", DB_CONNECTION)
	return writer
}
