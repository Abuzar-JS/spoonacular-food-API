package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/domain"
	"gorm.io/gorm"
)

type PreferencesPostgres struct {
	db *gorm.DB
}

func NewPreferencesPostgres(db *gorm.DB) *PreferencesPostgres {
	return &PreferencesPostgres{
		db: db,
	}
}

type IntoleranceRow struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

type IntoleranceRows []IntoleranceRow

func (u IntoleranceRow) TableName() string {
	return "intolerances"
}

func (u IntoleranceRow) toDomain() *domain.Intolerance {
	return &domain.Intolerance{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u IntoleranceRows) toDomain() domain.Intolerances {
	intolerances := make(domain.Intolerances, len(u))

	for i, allCuisines := range u {
		intolerances[i] = *allCuisines.toDomain()
	}
	return intolerances
}

type DietRow struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

type DietRows []DietRow

func (u DietRow) TableName() string {
	return "diets"
}

func (u DietRow) toDomain() *domain.Diet {
	return &domain.Diet{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u DietRows) toDomain() domain.Diets {
	diets := make(domain.Diets, len(u))

	for i, allDiets := range u {
		diets[i] = *allDiets.toDomain()
	}
	return diets
}

type CuisineRow struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

type CuisineRows []CuisineRow

func (u CuisineRow) TableName() string {
	return "cuisines"
}

func (u CuisineRow) toDomain() *domain.Cuisine {
	return &domain.Cuisine{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (u CuisineRows) toDomain() domain.Cuisines {
	cuisines := make(domain.Cuisines, len(u))

	for i, allCuisines := range u {
		cuisines[i] = *allCuisines.toDomain()
	}
	return cuisines
}

type UserIntoleranceRow struct {
	UserID        int `gorm:"column:user_id;primaryKey"`
	IntoleranceID int `gorm:"column:intolerance_id;primaryKey"`
}

type UserIntoleranceRows []UserIntoleranceRow

func (u UserIntoleranceRow) TableName() string {
	return "user_intolerances"
}

func (u UserIntoleranceRow) fromDomain(ud domain.UserIntolerance) UserIntoleranceRow {
	return UserIntoleranceRow{
		UserID:        ud.UserID,
		IntoleranceID: ud.IntoleranceID,
	}
}

func (u *UserIntoleranceRow) toDomain() *domain.UserIntolerance {
	return &domain.UserIntolerance{
		UserID:        u.UserID,
		IntoleranceID: u.IntoleranceID,
	}
}

func (u UserIntoleranceRows) toDomain() domain.UserIntolerances {
	intolerances := make(domain.UserIntolerances, len(u))
	for i, intolerance := range u {
		intolerances[i] = *intolerance.toDomain()
	}
	return intolerances
}

type UserDietRow struct {
	UserID int `gorm:"column:user_id;primaryKey"`
	DietID int `gorm:"column:diet_id;primaryKey"`
}

type UserDietRows []UserDietRow

func (u UserDietRow) TableName() string {
	return "user_diets"
}

func (u UserDietRow) fromDomain(ud domain.UserDiet) UserDietRow {
	return UserDietRow{
		UserID: ud.UserID,
		DietID: ud.DietID,
	}
}

func (u UserDietRow) toDomain() *domain.UserDiet {
	return &domain.UserDiet{
		UserID: u.UserID,
		DietID: u.DietID,
	}
}

func (u UserDietRows) toDomain() domain.UserDiets {
	diets := make(domain.UserDiets, len(u))
	for i, diet := range u {
		diets[i] = *diet.toDomain()
	}
	return diets
}

type UserCuisineRow struct {
	UserID    int `gorm:"column:user_id;primaryKey"`
	CuisineID int `gorm:"column:cuisine_id;primaryKey"`
}

type UserCuisineRows []UserCuisineRow

func (u UserCuisineRow) TableName() string {
	return "user_cuisines"
}

func (u UserCuisineRow) fromDomain(ud domain.UserCuisine) UserCuisineRow {
	return UserCuisineRow{
		UserID:    ud.UserID,
		CuisineID: ud.CuisineID,
	}
}

func (u UserCuisineRow) toDomain() *domain.UserCuisine {
	return &domain.UserCuisine{
		UserID:    u.UserID,
		CuisineID: u.CuisineID,
	}
}

func (u UserCuisineRows) toDomain() domain.UserCuisines {
	cuisines := make(domain.UserCuisines, len(u))
	for i, cuisine := range u {
		cuisines[i] = *cuisine.toDomain()
	}
	return cuisines
}

func (u PreferencesPostgres) GetIntolerances(ctx context.Context) (domain.Intolerances, error) {

	var intolerances IntoleranceRows

	if result := u.db.Order("id").Find(&intolerances); result.Error != nil {
		return nil, fmt.Errorf("failed to get intolerances: %w", result.Error)
	}

	return intolerances.toDomain(), nil
}

func (u PreferencesPostgres) GetDiets(ctx context.Context) (domain.Diets, error) {

	var diets DietRows

	if result := u.db.Order("id").Find(&diets); result.Error != nil {
		return nil, fmt.Errorf("failed to get diets: %w", result.Error)
	}

	return diets.toDomain(), nil
}

func (u PreferencesPostgres) GetCuisines(ctx context.Context) (domain.Cuisines, error) {

	var cuisines CuisineRows

	if result := u.db.Order("id").Find(&cuisines); result.Error != nil {
		return nil, fmt.Errorf("failed to get cuisines: %w", result.Error)
	}

	return cuisines.toDomain(), nil
}

func (u *PreferencesPostgres) SaveUserDiet(ctx context.Context, userID int, dietID int) (*domain.UserDiet, error) {
	diet := domain.UserDiet{
		UserID: userID,
		DietID: dietID,
	}

	dietRow := UserDietRow{}.fromDomain(diet)

	if err := u.db.Create(&dietRow).Error; err != nil {
		return nil, fmt.Errorf("failed to save user diets: %v", err)
	}

	return UserDietRow(dietRow).toDomain(), nil
}

func (u *PreferencesPostgres) SaveUserCuisine(ctx context.Context, userID int, cuisineID int) (*domain.UserCuisine, error) {
	cuisine := domain.UserCuisine{
		UserID:    userID,
		CuisineID: cuisineID,
	}
	cuisineRow := UserCuisineRow{}.fromDomain(cuisine)

	if err := u.db.Create(&cuisineRow).Error; err != nil {
		return nil, fmt.Errorf("failed to save user cuisine: %v", err)
	}

	return cuisineRow.toDomain(), nil
}

func (u *PreferencesPostgres) SaveUserIntolerance(ctx context.Context, userID int, intoleranceID int) (*domain.UserIntolerance, error) {
	intolerance := domain.UserIntolerance{
		UserID:        userID,
		IntoleranceID: intoleranceID,
	}
	intoleranceRow := UserIntoleranceRow{}.fromDomain(intolerance)

	if err := u.db.Create(&intoleranceRow).Error; err != nil {
		return nil, fmt.Errorf("failed to save user intolerance: %v", err)
	}

	return intoleranceRow.toDomain(), nil
}
