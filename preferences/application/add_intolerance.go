package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain/intolerance"
)

type CreateIntoleranceRequest struct {
	Name string
}

func (u CreateIntoleranceRequest) Validate(ctx context.Context) error {
	if u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	return nil

}

type CreateIntolerance func(ctx context.Context, request CreateIntoleranceRequest) (*domain.Intolerance, error)

func NewCreateIntolerance(
	IntoleranceRepo intolerance.Repository,
) CreateIntolerance {
	return func(ctx context.Context, request CreateIntoleranceRequest) (*domain.Intolerance, error) {
		if err := request.Validate(ctx); err != nil {
			return nil, err
		}

		intoleranceRequest := domain.Intolerance{
			Name: request.Name,
		}

		intoleranceCreated, err := IntoleranceRepo.Save(ctx, intoleranceRequest)
		if err != nil {
			return nil, fmt.Errorf("failed to create intolerance")
		}

		return &intoleranceCreated, nil
	}
}
