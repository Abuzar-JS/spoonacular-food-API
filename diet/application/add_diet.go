package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/diet/domain"
	diet "github.com/Abuzar-JS/go-spoonacular-api/diet/domain/diet"
)

type CreateDietRequest struct {
	Name string
}

func (u CreateDietRequest) Validate(ctx context.Context) error {
	if u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	return nil

}

type CreateDiet func(ctx context.Context, request CreateDietRequest) (*domain.Diet, error)

func NewCreateDiet(
	DietRepo diet.Repository,
) CreateDiet {
	return func(ctx context.Context, request CreateDietRequest) (*domain.Diet, error) {
		if err := request.Validate(ctx); err != nil {
			return nil, err
		}

		dietRequest := domain.Diet{
			Name: request.Name,
		}

		dietCreated, err := DietRepo.Save(ctx, dietRequest)
		if err != nil {
			return nil, fmt.Errorf("failed to create diet")
		}

		return &dietCreated, nil
	}
}
