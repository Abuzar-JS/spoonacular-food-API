package application

import (
	"github.com/Abuzar-JS/go-spoonacular-api/user/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/user/domain/user"
)

type GetUsers func() ([]domain.UserResponse, error)

func NewGetUsers(
	UserRepo user.Repository,
) GetUsers {
	return func() ([]domain.UserResponse, error) {

		result, err := UserRepo.GetAll()

		if err != nil {
			return nil, err
		}

		var AllUsers []domain.UserResponse

		for _, value := range result {
			User := domain.UserResponse{
				ID:   value.ID,
				Name: value.Name,
			}
			AllUsers = append(AllUsers, User)
		}
		return AllUsers, nil
	}
}
