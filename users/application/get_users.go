package application

import (
	"github.com/Abuzar-JS/go-spoonacular-api/users/domain"
	"github.com/Abuzar-JS/go-spoonacular-api/users/domain/user"
)

type GetUsers func() ([]domain.User, error)

func NewGetUsers(
	UserRepo user.Repository,
) GetUsers {
	return func() ([]domain.User, error) {

		result, err := UserRepo.GetAll()

		if err != nil {
			return nil, err
		}

		var AllUsers []domain.User

		for _, value := range result {
			User := domain.User{
				ID:       value.ID,
				Name:     value.Name,
				Location: value.Location,
				Password: value.Password,
			}
			AllUsers = append(AllUsers, User)
		}
		return AllUsers, nil
	}
}
