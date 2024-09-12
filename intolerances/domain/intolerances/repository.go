package intolerances

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/intolerances/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll() ([]domain.Intolerance, error)
	GetIntoleranceByID(intoleranceID int) (domain.Intolerance, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.Intolerance) (domain.Intolerance, error)
}
