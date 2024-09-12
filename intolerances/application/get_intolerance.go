package application

import (
	"github.com/Abuzar-JS/go-spoonacular-api/intolerances/domain"
	intolerance "github.com/Abuzar-JS/go-spoonacular-api/intolerances/domain/intolerances"
)

type GetIntolerances func() ([]domain.Intolerance, error)

func NewGetIntolerances(
	IntoleranceRepo intolerance.Repository,
) GetIntolerances {
	return func() ([]domain.Intolerance, error) {

		result, err := IntoleranceRepo.GetAll()

		if err != nil {
			return nil, err
		}

		var AllIntolerances []domain.Intolerance

		for _, value := range result {
			Intolerance := domain.Intolerance{
				ID:   value.ID,
				Name: value.Name,
			}
			AllIntolerances = append(AllIntolerances, Intolerance)
		}
		return AllIntolerances, nil
	}
}
