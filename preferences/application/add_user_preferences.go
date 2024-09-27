package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/preferences"
)

type AddUserPreferencesRequest struct {
	UserID       int
	Cuisines     []int
	Diets        []int
	Intolerances []int
}

type AddUserPreferencesResponse struct {
	UserCuisines     domain.UserCuisines
	UserDiets        domain.UserDiets
	UserIntolerances domain.UserIntolerances
}

type AddUserPreferences func(ctx context.Context, request AddUserPreferencesRequest) (*AddUserPreferencesResponse, error)

func NewAddUserPreferences(
	repo preferences.Repository,
) AddUserPreferences {
	return func(ctx context.Context, request AddUserPreferencesRequest) (*AddUserPreferencesResponse, error) {

		var userCuisines domain.UserCuisines
		for _, cuisineID := range request.Cuisines {
			userCuisine, err := repo.SaveUserCuisine(ctx, request.UserID, cuisineID)
			if err != nil {
				return nil, fmt.Errorf("failed to save user cuisine: %v", err)
			}
			userCuisines = append(userCuisines, *userCuisine)
		}

		var userDiets domain.UserDiets
		for _, dietID := range request.Diets {
			userDiet, err := repo.SaveUserDiet(ctx, request.UserID, dietID)
			if err != nil {
				return nil, fmt.Errorf("failed to save user diet: %v", err)
			}
			userDiets = append(userDiets, *userDiet)
		}

		var userIntolerances domain.UserIntolerances
		for _, intoleranceID := range request.Intolerances {
			userIntolerance, err := repo.SaveUserIntolerance(ctx, request.UserID, intoleranceID)
			if err != nil {
				return nil, fmt.Errorf("failed to save user intolerance: %v", err)
			}
			userIntolerances = append(userIntolerances, *userIntolerance)
		}

		return &AddUserPreferencesResponse{
			UserCuisines:     userCuisines,
			UserDiets:        userDiets,
			UserIntolerances: userIntolerances,
		}, nil
	}
}
