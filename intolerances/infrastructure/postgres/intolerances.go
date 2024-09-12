package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/intolerances/domain"

	"gorm.io/gorm"
)

type IntolerancePostgres struct {
	db *gorm.DB
}

func NewIntolerancePostgres(db *gorm.DB) *IntolerancePostgres {
	return &IntolerancePostgres{
		db: db,
	}
}

type Intolerance struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

func (u Intolerance) TableName() string {
	return "intolerances"
}

func (u Intolerance) ToDomain() domain.Intolerance {
	return domain.Intolerance{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u Intolerance) FromDomain(ud domain.Intolerance) Intolerance {
	return Intolerance{
		ID:   ud.ID,
		Name: ud.Name,
	}
}

func (u IntolerancePostgres) Save(ctx context.Context, request domain.Intolerance) (domain.Intolerance, error) {

	intolerance := Intolerance{}.FromDomain(request)

	result := u.db.Create(&intolerance)

	if result.Error != nil {
		return domain.Intolerance{}, fmt.Errorf("failed to create intolerance")
	}

	return intolerance.ToDomain(), nil

}

func (u IntolerancePostgres) GetIntoleranceByID(intoleranceID int) (domain.Intolerance, error) {

	var intolerance Intolerance

	result := u.db.First(&intolerance, intoleranceID)

	if result.Error != nil {
		return domain.Intolerance{}, fmt.Errorf("intolerance with ID %v not found", intoleranceID)
	}

	return intolerance.ToDomain(), nil
}

func (u IntolerancePostgres) GetAll() ([]domain.Intolerance, error) {
	var intolerances []domain.Intolerance

	result := u.db.Order("id").Find(&intolerances)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get intolerances: %w", result.Error)
	}

	return intolerances, nil
}
