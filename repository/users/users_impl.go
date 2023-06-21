package users

import (
	"context"
	"github.com/PyMarcus/go_crud/entities"
)

type UserRepositoryInterface interface {
	GetAll(ctx context.Context) ([]entities.User, error)
}
