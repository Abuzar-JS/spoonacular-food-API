package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/application"
	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {

	userCuisineRepo := postgres.NewUserCuisinePostgres(database)

	routes := router.Group("/api")

	routes.POST("user_cuisine/:user_id/:cuisine_id", NewAddUserCuisine(
		application.NewCreateUserCuisine(userCuisineRepo),
	))

	routes.GET("user_cuisine/:user_id", NewGetUserCuisines(
		application.NewGetCuisinesByUserID(userCuisineRepo),
	))

	return router
}
