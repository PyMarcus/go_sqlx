package users

import (
	"context"
	"github.com/PyMarcus/go_crud/entities"
	"github.com/jmoiron/sqlx"
	"log"
)

type userDbSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

// instancia um objecto com as conexoes de leitura e escrita seguindo o contrato com a interface
func NewSqlx(w *sqlx.DB, r *sqlx.DB) UserRepositoryInterface {
	return &userDbSqlx{writer: w, reader: r}
}

func (u *userDbSqlx) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User

	err := u.reader.SelectContext(ctx, &users, `SELECT * FROM users;`)

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return users, nil
}
