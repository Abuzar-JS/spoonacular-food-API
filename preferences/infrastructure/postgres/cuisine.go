package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
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

type CuisineRow struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

type CuisineRows []CuisineRow

func (u CuisineRow) TableName() string {
	return "cuisines"
}

func (u CuisineRow) ToDomain() *domain.Cuisine {
	return &domain.Cuisine{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u CuisineRows) ToDomain() domain.Cuisines {
	cuisines := make(domain.Cuisines, len(u))

	for i, allCuisines := range u {
		cuisines[i] = *allCuisines.ToDomain()
	}
	return cuisines
}

func (u CuisineRow) FromDomain(ud domain.Cuisine) CuisineRow {
	return CuisineRow{
		ID:   ud.ID,
		Name: ud.Name,
	}
}

func (u CuisinePostgres) Save(ctx context.Context, request domain.Cuisine) (*domain.Cuisine, error) {
	cuisine := CuisineRow{}.FromDomain(request)

	if result := u.db.Create(&cuisine); result.Error != nil {
		return nil, fmt.Errorf("failed to create cuisine")
	}

	return cuisine.ToDomain(), nil

}

func (u CuisinePostgres) GetByID(ctx context.Context, cuisineID int) (*domain.Cuisine, error) {
	var cuisine CuisineRow

	if result := u.db.First(&cuisine, cuisineID); result.Error != nil {
		return nil, fmt.Errorf("cuisine with ID %v not found", cuisineID)
	}

	return cuisine.ToDomain(), nil
}

func (u CuisinePostgres) Get(ctx context.Context) (domain.Cuisines, error) {

	var cuisines CuisineRows

	if result := u.db.Order("id").Find(&cuisines); result.Error != nil {
		return nil, fmt.Errorf("failed to get cuisines: %w", result.Error)
	}

	return cuisines.ToDomain(), nil
}
