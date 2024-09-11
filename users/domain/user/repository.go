package user

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/users/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll() ([]domain.User, error)
	GetUserByID(userID int) (User domain.User, err error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.User) (domain.User, error)
}
