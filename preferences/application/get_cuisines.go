package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/cuisine"
)

type GetCuisines func(ctx context.Context) (domain.Cuisines, error)

func NewGetCuisines(
	CuisineRepo cuisine.Repository,
) GetCuisines {
	return func(ctx context.Context) (domain.Cuisines, error) {

		allCuisines, err := CuisineRepo.Get(ctx)

		if err != nil {
			return nil, fmt.Errorf("failed to get cuisines")
		}

		return allCuisines, nil
	}
}
