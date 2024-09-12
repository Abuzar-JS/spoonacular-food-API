package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/users/application"
	"github.com/Abuzar-JS/go-spoonacular-api/users/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {
	userRepo := postgres.NewUserPostgres(database)

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
	

	return router
}
