package application

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/users/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/users/domain/user"
)

type GetUserByID func(ctx context.Context, userID int) (*domain.UserResponse, error)

func NewGetUserByID(
	userRepo user.Repository,
) GetUserByID {
	return func(ctx context.Context, userID int) (*domain.UserResponse, error) {

		user, err := userRepo.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
		userResponse := domain.UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Cuisine: user.Cuisine,
		}
		return &userResponse, nil
	}
}
