package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
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
	ID   int    `gorm:"column:id;primaryKey"`
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

	if result := u.db.Create(&intolerance); result.Error != nil {
		return domain.Intolerance{}, fmt.Errorf("failed to create intolerance")
	}

	return intolerance.ToDomain(), nil

}

func (u IntolerancePostgres) GetIntoleranceByID(ctx context.Context, intoleranceID int) (domain.Intolerance, error) {

	var intolerance Intolerance

	if result := u.db.First(&intolerance, intoleranceID); result.Error != nil {
		return domain.Intolerance{}, fmt.Errorf("intolerance with ID %v not found", intoleranceID)
	}

	return intolerance.ToDomain(), nil
}

func (u IntolerancePostgres) GetAll(ctx context.Context) ([]domain.Intolerance, error) {

	var intolerances []domain.Intolerance

	if result := u.db.Order("id").Find(&intolerances); result.Error != nil {
		return nil, fmt.Errorf("failed to get intolerances: %w", result.Error)
	}

	return intolerances, nil
}
