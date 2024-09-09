package application

import (
	"context"
	"fmt"

	"github.com/users/domain"
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

type CreateUser func(ctx context.Context, request CreateUserRequest) *domain.User
