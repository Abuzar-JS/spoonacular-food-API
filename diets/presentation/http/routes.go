package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/diets/application"
	"github.com/Abuzar-JS/go-spoonacular-api/diets/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {
	dietRepo := postgres.NewDietPostgres(database)

	dietRouter := router.Group("/api")

	dietRouter.POST("/diet", NewCreateDiet(
		application.NewCreateDiet(dietRepo),
	))

	dietRouter.GET("/diets", NewGetDiets(
		application.NewGetDiets(dietRepo),
	))

	dietRouter.GET("/diets/:diet_id", NewGetDietByID(
		application.NewGetDietByID(dietRepo),
	))

	return router
}
