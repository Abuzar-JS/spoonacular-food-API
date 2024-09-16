package diet

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll(ctx context.Context) ([]domain.Diet, error)
	GetDietByID(ctx context.Context, dietID int) (domain.Diet, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.Diet) (domain.Diet, error)
}
