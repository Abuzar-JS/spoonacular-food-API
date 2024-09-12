package application

import (
	"context"

	"github.com/Abuzar-JS/go-spoonacular-api/intolerances/domain"
	intolerance "github.com/Abuzar-JS/go-spoonacular-api/intolerances/domain/intolerances"
)

type GetIntoleranceByID func(ctx context.Context, intoleranceID int) (*domain.Intolerance, error)

func NewGetIntoleranceByID(
	intoleranceRepo intolerance.Repository,
) GetIntoleranceByID {
	return func(ctx context.Context, intoleranceID int) (*domain.Intolerance, error) {

		intolerance, err := intoleranceRepo.GetIntoleranceByID(intoleranceID)
		if err != nil {
			return nil, err
		}
		intoleranceResponse := domain.Intolerance{
			ID:   intolerance.ID,
			Name: intolerance.Name,
		}
		return &intoleranceResponse, nil
	}
}
