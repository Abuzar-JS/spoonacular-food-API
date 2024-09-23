package intolerance

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/intolerance/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll(ctx context.Context) ([]domain.Intolerance, error)
	GetIntoleranceByID(ctx context.Context, intoleranceID int) (domain.Intolerance, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.Intolerance) (domain.Intolerance, error)
}
