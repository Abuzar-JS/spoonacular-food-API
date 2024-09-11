package postgres

import (
	"context"
	"fmt"

	"github.com/Abuzar-JS/spoonacular-food-API/users/domain"

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
	Location string `gorm:"column:location"`
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
		Location: u.Location,
		Password: u.Password,
	}
}

func (u User) FromDomain(ud domain.User) User {
	return User{
		ID:       ud.ID,
		Name:     ud.Name,
		Location: ud.Location,
		Password: ud.Password,
	}
}

type UserChoice struct {
	UserID        int `gorm:"primaryKey"`
	IngredientsID int `gorm:"primaryKey"`
	IsUserChoice  bool
	User          User        `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Ingredient    Ingredients `gorm:"foreignKey:IngredientsID;references:ID;constraint:OnDelete:CASCADE"`
}

func (u UserChoice) TableName() string {
	return "user_choices"
}

type Ingredients struct {
	ID int `gorm:"primaryKey"`
}

type UserMeals struct {
	UserID   int    `gorm:"primaryKey"`
	RecipeID int    `gorm:"primaryKey"`
	MealType string `gorm:"size:50"`
	User     User   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Recipe   Recipe `gorm:"foreignKey:RecipeID;references:ID;constraint:OnDelete:CASCADE"`
}

func (u UserMeals) TableName() string {
	return "user_meals"
}

type Recipe struct {
	ID int `gorm:"primaryKey"`
}

func (u UserPostgres) Save(ctx context.Context, request domain.User) (domain.User, error) {

	user := User{}.FromDomain(request)

	result := u.db.Create(&user)

	if result != nil {
		return domain.User{}, fmt.Errorf("failed to create user")
	}

	return user.ToDomain(), nil

}
