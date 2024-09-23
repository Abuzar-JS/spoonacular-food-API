package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/user/application"
	"github.com/Abuzar-JS/go-spoonacular-api/user/infrastructure/postgres"
	"github.com/Abuzar-JS/go-spoonacular-api/user/infrastructure/spoonacular"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {
	userRepo := postgres.NewUserPostgres(database)

	recipeRepo := &spoonacular.SpoonacularClient{}

	userRouter := router.Group("/api")

	//Create User

	userRouter.POST("/user", NewCreateUser(
		application.NewCreateUser(userRepo),
	))

	userRouter.GET("/users", NewGetUsers(
		application.NewGetUsers(userRepo),
	))

	userRouter.GET("/users/:user_id", NewGetUserByID(
		application.NewGetUserByID(userRepo),
	))

	userRouter.GET("/recipes", NewSpoonacularRecipe(
		application.NewSpoonacularRecipe(recipeRepo),
	))

	return router
}
