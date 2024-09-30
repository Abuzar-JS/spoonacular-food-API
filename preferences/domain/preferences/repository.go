package preferences

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"gorm.io/gorm"
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
	SaveUserDiet(ctx context.Context, tx *gorm.DB, userID int, dietID int) (*domain.UserDiet, error)
	SaveUserCuisine(ctx context.Context, tx *gorm.DB, userID int, cuisineID int) (*domain.UserCuisine, error)
	SaveUserIntolerance(ctx context.Context, tx *gorm.DB, userID int, intoleranceID int) (*domain.UserIntolerance, error)
	StartTransaction() (*gorm.DB, error)
	// SaveUserPreferences(ctx context.Context, userID, dietID, cuisineID, intoleranceID int) error
}
