package application

import (
	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain"
	cuisine "github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain/cuisines"
)

type GetCuisines func() ([]domain.Cuisine, error)

func NewGetCuisines(
	CuisineRepo cuisine.Repository,
) GetCuisines {
	return func() ([]domain.Cuisine, error) {

		result, err := CuisineRepo.GetAll()

		if err != nil {
			return nil, err
		}

		var AllCuisines []domain.Cuisine

		for _, value := range result {
			Cuisine := domain.Cuisine{
				ID:   value.ID,
				Name: value.Name,
			}
			AllCuisines = append(AllCuisines, Cuisine)
		}
		return AllCuisines, nil
	}
}
