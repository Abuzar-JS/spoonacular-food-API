package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisine/domain"
	cuisine "github.com/Abuzar-JS/go-spoonacular-api/cuisine/domain/cuisine"
)

type GetCuisineByID func(context.Context, int) (*domain.Cuisine, error)

func NewGetCuisineByID(
	cuisineRepo cuisine.Repository,
) GetCuisineByID {
	return func(ctx context.Context, cuisineID int) (*domain.Cuisine, error) {

		cuisine, err := cuisineRepo.GetByID(ctx, cuisineID)
		if err != nil {
			return nil, fmt.Errorf("cuisine with ID %v not found", cuisineID)
		}
		cuisineResponse := domain.Cuisine{
			ID:   cuisine.ID,
			Name: cuisine.Name,
		}
		return &cuisineResponse, nil
	}
}
