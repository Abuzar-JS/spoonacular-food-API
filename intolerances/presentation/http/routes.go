package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/intolerances/application"
	"github.com/Abuzar-JS/go-spoonacular-api/intolerances/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {
	intoleranceRepo := postgres.NewIntolerancePostgres(database)

	intoleranceRouter := router.Group("/api")

	intoleranceRouter.POST("/intolerance", NewCreateIntolerance(
		application.NewCreateIntolerance(intoleranceRepo),
	))

	intoleranceRouter.GET("/intolerances", NewGetIntolerances(
		application.NewGetIntolerances(intoleranceRepo),
	))

	intoleranceRouter.GET("/intolerances/:intolerance_id", NewGetIntoleranceByID(
		application.NewGetIntoleranceByID(intoleranceRepo),
	))

	return router
}
