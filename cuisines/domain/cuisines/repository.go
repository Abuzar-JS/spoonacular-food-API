package cuisines

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll() ([]domain.Cuisine, error)
	GetCuisineByID(cuisineID int) (domain.Cuisine, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.Cuisine) (domain.Cuisine, error)
}
