package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/diet"
)

type GetDietByID func(context.Context, int) (*domain.Diet, error)

func NewGetDietByID(
	dietRepo diet.Repository,
) GetDietByID {
	return func(ctx context.Context, dietID int) (*domain.Diet, error) {

		diet, err := dietRepo.GetDietByID(ctx, dietID)
		if err != nil {
			return nil, fmt.Errorf("diet with ID %v not found", dietID)
		}
		dietResponse := domain.Diet{
			ID:   diet.ID,
			Name: diet.Name,
		}
		return &dietResponse, nil
	}
}
