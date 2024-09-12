package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisine/domain"
	cuisine "github.com/Abuzar-JS/go-spoonacular-api/cuisine/domain/cuisine"
)

type GetCuisines func(ctx context.Context) ([]domain.Cuisine, error)

func NewGetCuisines(
	CuisineRepo cuisine.Repository,
) GetCuisines {
	return func(ctx context.Context) ([]domain.Cuisine, error) {

		result, err := CuisineRepo.GetAll(ctx)

		if err != nil {
			return nil, fmt.Errorf("failed to get cuisines")
		}

		var allCuisines []domain.Cuisine

		for _, value := range result {
			Cuisine := domain.Cuisine{
				ID:   value.ID,
				Name: value.Name,
			}
			allCuisines = append(allCuisines, Cuisine)
		}
		return allCuisines, nil
	}
}
