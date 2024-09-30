package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/application"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {

	preferencesRepo := postgres.NewPreferencesPostgres(database)

	routes := router.Group("/api")

	// routes.POST("/user_preferences", NewAddUserPreferences(
	// 	application.NewAddUserPreferences(preferencesRepo),
	// ))

	routes.POST("/user_preferences", NewAddUserPreferences(
		application.NewAddUserPreferences(preferencesRepo),
	))

	routes.GET("/intolerances", NewGetIntolerances(
		application.NewGetIntolerances(preferencesRepo),
	))

	routes.GET("/diets", NewGetDiets(
		application.NewGetDiets(preferencesRepo),
	))

	routes.GET("/cuisines", NewGetCuisines(
		application.NewGetCuisines(preferencesRepo),
	))

	return router
}
