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
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.User) (domain.User, error)
}
