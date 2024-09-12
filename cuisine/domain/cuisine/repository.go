package cuisine

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisine/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll(ctx context.Context) ([]domain.Cuisine, error)
	GetCuisineByID(ctx context.Context, cuisineID int) (domain.Cuisine, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.Cuisine) (domain.Cuisine, error)
}
