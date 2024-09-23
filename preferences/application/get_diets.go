package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/diet"
)

type GetDiets func(ctx context.Context) ([]domain.Diet, error)

func NewGetDiets(
	DietRepo diet.Repository,
) GetDiets {
	return func(ctx context.Context) ([]domain.Diet, error) {

		result, err := DietRepo.GetAll(ctx)

		if err != nil {
			return nil, fmt.Errorf("failed to get diets")
		}

		var allDiets []domain.Diet

		for _, value := range result {
			Diet := domain.Diet{
				ID:   value.ID,
				Name: value.Name,
			}
			allDiets = append(allDiets, Diet)
		}
		return allDiets, nil
	}
}
