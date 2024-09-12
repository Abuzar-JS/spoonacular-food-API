package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/cuisine/application"
	"github.com/Abuzar-JS/go-spoonacular-api/cuisine/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {
	cuisineRepo := postgres.NewCuisinePostgres(database)

	cuisineRouter := router.Group("/api")

	cuisineRouter.POST("/cuisine", NewCreateCuisine(
		application.NewCreateCuisine(cuisineRepo),
	))

	cuisineRouter.GET("/cuisines", NewGetCuisines(
		application.NewGetCuisines(cuisineRepo),
	))

	cuisineRouter.GET("/cuisines/:cuisine_id", NewGetCuisineByID(
		application.NewGetCuisineByID(cuisineRepo),
	))

	return router
}
