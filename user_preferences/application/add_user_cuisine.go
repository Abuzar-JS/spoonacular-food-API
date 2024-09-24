package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/domain"
	cuisine "github.com/Abuzar-JS/go-spoonacular-api/user_preferences/domain/user_cuisine"
)

type CreateUserCuisineRequest struct {
	UserID    int
	CuisineID int
}

func (u CreateUserCuisineRequest) Validate(ctx context.Context) error {
	if u.UserID <= 0 {
		return fmt.Errorf("user ID must be greater than zero")
	}
	if u.CuisineID <= 0 {
		return fmt.Errorf("cuisine ID must be greater than zero")
	}
	return nil
}

type CreateUserCuisine func(ctx context.Context, request CreateUserCuisineRequest) (*domain.UserCuisine, error)

func NewCreateUserCuisine(
	userCuisineRepo cuisine.Repository,
) CreateUserCuisine {
	return func(ctx context.Context, request CreateUserCuisineRequest) (*domain.UserCuisine, error) {
		if err := request.Validate(ctx); err != nil {
			return nil, err
		}

		userCuisineRequest := domain.UserCuisine{
			UserID:    request.UserID,
			CuisineID: request.CuisineID,
		}

		userCuisineCreated, err := userCuisineRepo.Save(ctx, userCuisineRequest)
		if err != nil {
			return nil, fmt.Errorf("failed to create user cuisine: %w", err)
		}

		return userCuisineCreated, nil
	}
}
