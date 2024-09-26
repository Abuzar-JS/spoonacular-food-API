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
	GetIntolerances(ctx context.Context) ([]domain.Intolerance, error)
	GetDiets(ctx context.Context) ([]domain.Diet, error)
}

type WriteRepository interface {
	SaveUserCuisines(ctx context.Context, cuisines domain.UserCuisines) (domain.UserCuisines, error)
	SaveUserDiets(ctx context.Context, diets domain.UserDiets) (domain.UserDiets, error)
	SaveUserIntolerances(ctx context.Context, intolerances domain.UserIntolerances) (domain.UserIntolerances, error)
}
