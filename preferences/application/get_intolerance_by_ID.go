package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/intolerance"
)

type GetIntoleranceByID func(context.Context, int) (*domain.Intolerance, error)

func NewGetIntoleranceByID(
	intoleranceRepo intolerance.Repository,
) GetIntoleranceByID {
	return func(ctx context.Context, intoleranceID int) (*domain.Intolerance, error) {

		intolerance, err := intoleranceRepo.GetIntoleranceByID(ctx, intoleranceID)
		if err != nil {
			return nil, fmt.Errorf("intolerance with ID %v not found", intoleranceID)
		}
		intoleranceResponse := domain.Intolerance{
			ID:   intolerance.ID,
			Name: intolerance.Name,
		}
		return &intoleranceResponse, nil
	}
}
