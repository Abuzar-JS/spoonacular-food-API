package preferences

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetCuisines(ctx context.Context) (domain.Cuisines, error)
	GetIntolerances(ctx context.Context) (domain.Intolerances, error)
	GetDiets(ctx context.Context) (domain.Diets, error)
}

type WriteRepository interface {
	SaveUserPreferences(ctx context.Context, userID int, cuisineIDs, dietIDs, intoleranceIDs []int) (*domain.UserPreferences, error)
}
