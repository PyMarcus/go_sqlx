package repository

import (
	"github.com/PyMarcus/go_crud/repository/users"
	"github.com/jmoiron/sqlx"
)

type Option struct {
	WriterSqlx *sqlx.DB
	ReadSqlx   *sqlx.DB
}

type Container struct {
	User users.UserRepositoryInterface
}

func New(options Option) *Container {
	return &Container{User: users.NewSqlx(options.WriterSqlx, options.ReadSqlx)}
}
