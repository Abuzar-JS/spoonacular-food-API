package application

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/diets/domain"
	diet "github.com/Abuzar-JS/go-spoonacular-api/diets/domain/diets"
)

type GetDietByID func(ctx context.Context, dietID int) (*domain.Diet, error)

func NewGetDietByID(
	dietRepo diet.Repository,
) GetDietByID {
	return func(ctx context.Context, dietID int) (*domain.Diet, error) {

		diet, err := dietRepo.GetDietByID(dietID)
		if err != nil {
			return nil, err
		}
		dietResponse := domain.Diet{
			ID:   diet.ID,
			Name: diet.Name,
		}
		return &dietResponse, nil
	}
}
