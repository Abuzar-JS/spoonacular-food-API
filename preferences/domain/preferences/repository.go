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
	SaveUserDiet(ctx context.Context, userID int, dietID int) (*domain.UserDiet, error)
	SaveUserCuisine(ctx context.Context, userID int, cuisineID int) (*domain.UserCuisine, error)
	SaveUserIntolerance(ctx context.Context, userID int, intoleranceID int) (*domain.UserIntolerance, error)
}
