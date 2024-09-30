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

		// Start a transaction
		tx, err := repo.StartTransaction()
		if err != nil {
			return nil, fmt.Errorf("failed to start transaction: %v", err)
		}

		// Defer a function to handle transaction commit or rollback
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r) // re-throw panic after rollback
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit().Error
				if err != nil {
					tx.Rollback()
				}
			}
		}()

		var userCuisines domain.UserCuisines
		for _, cuisineID := range request.Cuisines {
			userCuisine, err := repo.SaveUserCuisine(ctx, tx, request.UserID, cuisineID)
			if err != nil {
				return nil, fmt.Errorf("failed to save user cuisine: %v", err)
			}
			userCuisines = append(userCuisines, *userCuisine)
		}

		var userDiets domain.UserDiets
		for _, dietID := range request.Diets {
			userDiet, err := repo.SaveUserDiet(ctx, tx, request.UserID, dietID)
			if err != nil {
				return nil, fmt.Errorf("failed to save user diet: %v", err)
			}
			userDiets = append(userDiets, *userDiet)
		}

		var userIntolerances domain.UserIntolerances
		for _, intoleranceID := range request.Intolerances {
			userIntolerance, err := repo.SaveUserIntolerance(ctx, tx, request.UserID, intoleranceID)
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
