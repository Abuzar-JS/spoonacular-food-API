package diet

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/diets/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetAll() ([]domain.Diet, error)
	GetDietByID(dietID int) (domain.Diet, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.Diet) (domain.Diet, error)
}
