package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/domain"

	"gorm.io/gorm"
)

type CuisinePostgres struct {
	db *gorm.DB
}

func NewCuisinePostgres(db *gorm.DB) *CuisinePostgres {
	return &CuisinePostgres{
		db: db,
	}
}

type Cuisine struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

func (u Cuisine) TableName() string {
	return "cuisines"
}

func (u Cuisine) ToDomain() domain.Cuisine {
	return domain.Cuisine{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u Cuisine) FromDomain(ud domain.Cuisine) Cuisine {
	return Cuisine{
		ID:   ud.ID,
		Name: ud.Name,
	}
}

func (u CuisinePostgres) Save(ctx context.Context, request domain.Cuisine) (domain.Cuisine, error) {

	cuisine := Cuisine{}.FromDomain(request)

	result := u.db.Create(&cuisine)

	if result.Error != nil {
		return domain.Cuisine{}, fmt.Errorf("failed to create cuisine")
	}

	return cuisine.ToDomain(), nil

}

func (u CuisinePostgres) GetCuisineByID(cuisineID int) (domain.Cuisine, error) {

	var cuisine Cuisine

	result := u.db.First(&cuisine, cuisineID)

	if result.Error != nil {
		return domain.Cuisine{}, fmt.Errorf("cuisine with ID %v not found", cuisineID)
	}

	return cuisine.ToDomain(), nil
}

func (u CuisinePostgres) GetAll() ([]domain.Cuisine, error) {
	var cuisines []domain.Cuisine

	result := u.db.Order("id").Find(&cuisines)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get cuisines: %w", result.Error)
	}

	return cuisines, nil
}
