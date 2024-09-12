package application

import (
	"github.com/Abuzar-JS/go-spoonacular-api/diets/domain"
	diet "github.com/Abuzar-JS/go-spoonacular-api/diets/domain/diets"
)

type GetDiets func() ([]domain.Diet, error)

func NewGetDiets(
	DietRepo diet.Repository,
) GetDiets {
	return func() ([]domain.Diet, error) {

		result, err := DietRepo.GetAll()

		if err != nil {
			return nil, err
		}

		var AllDiets []domain.Diet

		for _, value := range result {
			Diet := domain.Diet{
				ID:   value.ID,
				Name: value.Name,
			}
			AllDiets = append(AllDiets, Diet)
		}
		return AllDiets, nil
	}
}
