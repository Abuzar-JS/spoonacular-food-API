package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisine/domain"

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
	ID   int    `gorm:"column:id;primaryKey"`
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

	if result := u.db.Create(&cuisine); result.Error != nil {
		return domain.Cuisine{}, fmt.Errorf("failed to create cuisine")
	}

	return cuisine.ToDomain(), nil

}

func (u CuisinePostgres) GetCuisineByID(ctx context.Context, cuisineID int) (domain.Cuisine, error) {

	var cuisine Cuisine

	if result := u.db.First(&cuisine, cuisineID); result.Error != nil {
		return domain.Cuisine{}, fmt.Errorf("cuisine with ID %v not found", cuisineID)
	}

	return cuisine.ToDomain(), nil
}

func (u CuisinePostgres) GetAll(ctx context.Context) ([]domain.Cuisine, error) {

	var cuisines []domain.Cuisine

	if result := u.db.Order("id").Find(&cuisines); result.Error != nil {
		return nil, fmt.Errorf("failed to get cuisines: %w", result.Error)
	}

	return cuisines, nil
}
