package application

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain"
	cuisine "github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain/cuisines"
)

type GetCuisineByID func(ctx context.Context, cuisineID int) (*domain.Cuisine, error)

func NewGetCuisineByID(
	cuisineRepo cuisine.Repository,
) GetCuisineByID {
	return func(ctx context.Context, cuisineID int) (*domain.Cuisine, error) {

		cuisine, err := cuisineRepo.GetCuisineByID(cuisineID)
		if err != nil {
			return nil, err
		}
		cuisineResponse := domain.Cuisine{
			ID:   cuisine.ID,
			Name: cuisine.Name,
		}
		return &cuisineResponse, nil
	}
}
