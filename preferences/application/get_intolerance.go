package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/intolerance"
)

type GetIntolerances func(ctx context.Context) ([]domain.Intolerance, error)

func NewGetIntolerances(
	IntoleranceRepo intolerance.Repository,
) GetIntolerances {
	return func(ctx context.Context) ([]domain.Intolerance, error) {

		result, err := IntoleranceRepo.GetAll(ctx)

		if err != nil {
			return nil, fmt.Errorf("failed to get intolerances")
		}

		var allIntolerances []domain.Intolerance

		for _, value := range result {
			Intolerance := domain.Intolerance{
				ID:   value.ID,
				Name: value.Name,
			}
			allIntolerances = append(allIntolerances, Intolerance)
		}
		return allIntolerances, nil
	}
}
