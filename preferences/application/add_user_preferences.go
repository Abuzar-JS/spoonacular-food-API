package application

import (
	"context"

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
	UserPreferences domain.UserPreferences
}

type AddUserPreferences func(ctx context.Context, request AddUserPreferencesRequest) (*AddUserPreferencesResponse, error)

func NewAddUserPreferences(
	repo preferences.Repository,
) AddUserPreferences {
	return func(ctx context.Context, request AddUserPreferencesRequest) (*AddUserPreferencesResponse, error) {
		userPreferences, err := repo.SaveUserPreferences(ctx, request.UserID, request.Cuisines, request.Diets, request.Intolerances)
		if err != nil {
			return nil, err
		}

		return &AddUserPreferencesResponse{
			UserPreferences: *userPreferences,
		}, nil
	}
}
