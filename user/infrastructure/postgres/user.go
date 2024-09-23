package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/go-spoonacular-api/user/domain"

	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:name;not null;unique"`
	Password string `gorm:"column:password"`
}

func (u User) TableName() string {
	return "users"
}

type Users []User

func (u User) ToDomain() domain.User {
	return domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
	}
}

func (u User) FromDomain(ud domain.User) User {
	return User{
		ID:       ud.ID,
		Name:     ud.Name,
		Password: ud.Password,
	}
}

func (u UserPostgres) Save(ctx context.Context, request domain.User) (domain.User, error) {

	user := User{}.FromDomain(request)

	result := u.db.Create(&user)

	if result.Error != nil {
		return domain.User{}, fmt.Errorf("failed to create user")
	}

	return user.ToDomain(), nil

}

func (u UserPostgres) GetUserByID(userID int) (domain.User, error) {

	var user User

	result := u.db.First(&user, userID)

	if result.Error != nil {
		return domain.User{}, fmt.Errorf("user with ID %v not found", userID)
	}

	return user.ToDomain(), nil
}

func (u UserPostgres) GetAll() ([]domain.User, error) {
	var users []domain.User

	result := u.db.Order("id").Find(&users)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get users: %w", result.Error)
	}

	return users, nil
}
