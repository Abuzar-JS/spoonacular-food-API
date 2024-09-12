package user

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/user/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll() ([]domain.User, error)
	GetUserByID(userID int) (domain.User, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.User) (domain.User, error)
}
