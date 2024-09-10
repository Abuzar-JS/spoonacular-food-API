package application

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/pkg/helper"
	"github.com/Abuzar-JS/go-spoonacular-api/users/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/users/domain/user"
)

type CreateUserRequest struct {
	Name     string
	Location string
	Password string
}

func (u CreateUserRequest) Validate(ctx context.Context) error {
	if u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if u.Location == "" {
		return fmt.Errorf("location cannot be empty")
	}

	if u.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	return nil

}

type CreateUser func(ctx context.Context, request CreateUserRequest) (*domain.User, error)

func NewCreateUser(
	UserRepo user.Repository,
) CreateUser {
	return func(ctx context.Context, request CreateUserRequest) (*domain.User, error) {
		if err := request.Validate(ctx); err != nil {
			return nil, err
		}

		hashedPassword, err := helper.HashPassword(request.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}

		userRequest := domain.User{
			Name:     request.Name,
			Location: request.Location,
			Password: hashedPassword,
		}

		userCreated, err := UserRepo.Save(ctx, userRequest)
		if err != nil {
			return nil, fmt.Errorf("user creation failed")

		}
		return &userCreated, nil
	}
}
