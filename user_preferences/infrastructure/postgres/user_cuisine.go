package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/domain"
	"gorm.io/gorm"
)

type UserCuisinePostgres struct {
	db *gorm.DB
}

func NewUserCuisinePostgres(db *gorm.DB) *UserCuisinePostgres {
	return &UserCuisinePostgres{
		db: db,
	}
}

type User struct {
	ID int `gorm:"column:id"`
}

type UserCuisineRow struct {
	UserID    int     `gorm:"column:user_id;primaryKey"`
	CuisineID int     `gorm:"column:cuisine_id;primaryKey"`
	Cuisine   Cuisine `gorm:"foreignKey:CuisineID;references:ID"`
}

type Cuisine struct {
	ID   int    `gorm:"column:id;"`
	Name string `gorm:"column:name;"`
}

type UserCuisineRows []UserCuisineRow

func (u UserCuisineRow) TableName() string {
	return "user_cuisines"
}

func (u UserCuisineRow) ToDomain() *domain.UserCuisine {
	return &domain.UserCuisine{
		UserID:    u.UserID,
		CuisineID: u.CuisineID,
		Cuisine: domain.Cuisine{
			Name: u.Cuisine.Name,
		},
	}
}

func (u UserCuisineRows) ToDomain() domain.UserCuisines {
	userCuisines := make(domain.UserCuisines, len(u))

	for i, allUserCuisines := range u {
		userCuisines[i] = *allUserCuisines.ToDomain()
	}
	return userCuisines
}

func (u UserCuisineRow) FromDomain(ud domain.UserCuisine) UserCuisineRow {
	return UserCuisineRow{
		UserID:    ud.UserID,
		CuisineID: ud.CuisineID,
	}
}

func (u *UserCuisinePostgres) Save(ctx context.Context, request domain.UserCuisine) (*domain.UserCuisine, error) {

	// var user User

	// result := u.db.First(&user, request.UserID)

	// if result.Error != nil {
	// 	return nil, fmt.Errorf("user with ID %v not found", request.UserID)
	// }

	// var cuisine Cuisine

	// result = u.db.First(&cuisine, request.CuisineID)

	// if result.Error != nil {
	// 	return nil, fmt.Errorf("cuisine with ID %v not found", request.CuisineID)
	// }

	userCuisine := UserCuisineRow{}.FromDomain(request)

	if result := u.db.Create(&userCuisine); result.Error != nil {
		return nil, fmt.Errorf("failed to add user cuisine")
	}

	return userCuisine.ToDomain(), nil

}

func (u *UserCuisinePostgres) GetCuisinesByUserID(ctx context.Context, userID int) (domain.UserCuisines, error) {

	user := UserCuisineRow{UserID: userID}

	result := u.db.First(&user, userID)

	if result.Error != nil {
		return nil, fmt.Errorf("user with ID %v not found", userID)
	}

	var userCuisines []UserCuisineRow

	err := u.db.Preload("Cuisine").
		Where("user_id = ?", userID).
		Find(&userCuisines).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get user cuisines: %v", err)
	}

	return UserCuisineRows(userCuisines).ToDomain(), nil
}
