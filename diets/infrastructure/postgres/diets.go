package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/diets/domain"

	"gorm.io/gorm"
)

type DietPostgres struct {
	db *gorm.DB
}

func NewDietPostgres(db *gorm.DB) *DietPostgres {
	return &DietPostgres{
		db: db,
	}
}

type Diet struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

func (u Diet) TableName() string {
	return "diets"
}

func (u Diet) ToDomain() domain.Diet {
	return domain.Diet{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u Diet) FromDomain(ud domain.Diet) Diet {
	return Diet{
		ID:   ud.ID,
		Name: ud.Name,
	}
}

func (u DietPostgres) Save(ctx context.Context, request domain.Diet) (domain.Diet, error) {

	diet := Diet{}.FromDomain(request)

	result := u.db.Create(&diet)

	if result.Error != nil {
		return domain.Diet{}, fmt.Errorf("failed to create diet")
	}

	return diet.ToDomain(), nil

}

func (u DietPostgres) GetDietByID(dietID int) (domain.Diet, error) {

	var diet Diet

	result := u.db.First(&diet, dietID)

	if result.Error != nil {
		return domain.Diet{}, fmt.Errorf("diet with ID %v not found", dietID)
	}

	return diet.ToDomain(), nil
}

func (u DietPostgres) GetAll() ([]domain.Diet, error) {
	var diets []domain.Diet

	result := u.db.Order("id").Find(&diets)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get diets: %w", result.Error)
	}

	return diets, nil
}
