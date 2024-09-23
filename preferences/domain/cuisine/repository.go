package cuisine

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	Get(ctx context.Context) (domain.Cuisines, error)
	GetByID(ctx context.Context, cuisineID int) (*domain.Cuisine, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.Cuisine) (*domain.Cuisine, error)
}
