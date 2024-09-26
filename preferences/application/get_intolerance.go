package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/preferences"
)

type GetIntolerances func(ctx context.Context) ([]domain.Intolerance, error)

func NewGetIntolerances(
	repo preferences.Repository,
) GetIntolerances {
	return func(ctx context.Context) ([]domain.Intolerance, error) {

		result, err := repo.GetIntolerances(ctx)

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
