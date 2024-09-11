package application

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/users/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/users/domain/user"
)

type GetUserByID func(ctx context.Context, userID int) (*domain.User, error)

func NewGetUserByID(
	userRepo user.Repository,
) GetUserByID {
	return func(ctx context.Context, userID int) (*domain.User, error) {

		User, err := userRepo.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
		userResponse := domain.User{
			ID:      User.ID,
			Name:    User.Name,
			Cuisine: User.Cuisine,
		}
		return &userResponse, nil
	}
}
