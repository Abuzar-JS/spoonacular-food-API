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

func (u *UserIntoleranceRows) toDomain() domain.UserIntolerances {
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

func (u *PreferencesPostgres) SaveUserIntolerances(ctx context.Context, intolerances domain.UserIntolerances) (domain.UserIntolerances, error) {
	var intoleranceRows UserIntoleranceRows

	for _, intolerance := range intolerances {
		intoleranceRow := UserIntoleranceRow{}.fromDomain(intolerance)
		intoleranceRows = append(intoleranceRows, intoleranceRow)
	}

	if err := u.db.Create(&intoleranceRows).Error; err != nil {
		return nil, fmt.Errorf("failed to save user intolerances: %v", err)
	}

	return UserIntoleranceRows(intoleranceRows).toDomain(), nil
}

func (u *PreferencesPostgres) SaveUserDiets(ctx context.Context, diets domain.UserDiets) (domain.UserDiets, error) {
	var dietRows UserDietRows

	for _, diet := range diets {
		dietRow := UserDietRow{}.fromDomain(diet)
		dietRows = append(dietRows, dietRow)
	}

	if err := u.db.Create(&dietRows).Error; err != nil {
		return nil, fmt.Errorf("failed to save user diets: %v", err)
	}

	return UserDietRows(dietRows).toDomain(), nil
}

func (u *PreferencesPostgres) SaveUserCuisines(ctx context.Context, cuisines domain.UserCuisines) (domain.UserCuisines, error) {
	var cuisineRows UserCuisineRows

	for _, cuisine := range cuisines {
		cuisineRow := UserCuisineRow{}.fromDomain(cuisine)
		cuisineRows = append(cuisineRows, cuisineRow)
	}

	if err := u.db.Create(&cuisineRows).Error; err != nil {
		return nil, fmt.Errorf("failed to save user cuisines: %v", err)
	}

	return UserCuisineRows(cuisineRows).toDomain(), nil
}

func (u *PreferencesPostgres) SaveUserIntolerance(ctx context.Context, userID int, intoleranceID int) (*domain.UserIntolerances, error) {
	intoleranceRow := UserIntoleranceRow{
		UserID:        userID,
		IntoleranceID: intoleranceID,
	}

	if err := u.db.Create(&intoleranceRow).Error; err != nil {
		return nil, fmt.Errorf("failed to save user intolerance: %v", err)
	}
	return into, nil
}
