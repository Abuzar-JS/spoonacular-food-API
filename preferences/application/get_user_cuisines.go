package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/domain"
	cuisine "github.com/Abuzar-JS/go-spoonacular-api/user_preferences/domain/user_cuisine"
)

type GetCuisinesByUserIDRequest struct {
	UserID int
}

type GetCuisinesByUserID func(ctx context.Context, request GetCuisinesByUserIDRequest) (domain.UserCuisines, error)

func NewGetCuisinesByUserID(
	userCuisineRepo cuisine.Repository,
) GetCuisinesByUserID {
	return func(ctx context.Context, request GetCuisinesByUserIDRequest) (domain.UserCuisines, error) {
		if request.UserID <= 0 {
			return nil, fmt.Errorf("user ID must be greater than zero")
		}

		userCuisines, err := userCuisineRepo.GetCuisinesByUserID(ctx, request.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to get cuisines for user: %w", err)
		}

		return userCuisines, nil
	}
}
