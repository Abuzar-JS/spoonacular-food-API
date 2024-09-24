package cuisine

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	GetCuisinesByUserID(ctx context.Context, userID int) (domain.UserCuisines, error)
}

type WriteRepository interface {
	Save(ctx context.Context, request domain.UserCuisine) (*domain.UserCuisine, error)
}
