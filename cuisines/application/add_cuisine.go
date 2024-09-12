package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain"
	cuisine "github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain/cuisines"
)

type CreateCuisineRequest struct {
	Name string
}

func (u CreateCuisineRequest) Validate(ctx context.Context) error {
	if u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	return nil

}

type CreateCuisine func(ctx context.Context, request CreateCuisineRequest) (*domain.Cuisine, error)

func NewCreateCuisine(
	CuisineRepo cuisine.Repository,
) CreateCuisine {
	return func(ctx context.Context, request CreateCuisineRequest) (*domain.Cuisine, error) {
		if err := request.Validate(ctx); err != nil {
			return nil, err
		}

		cuisineRequest := domain.Cuisine{
			Name: request.Name,
		}

		cuisineCreated, err := CuisineRepo.Save(ctx, cuisineRequest)
		if err != nil {
			return nil, err
		}

		return &cuisineCreated, nil
	}
}
