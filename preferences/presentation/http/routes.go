package http

import (
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/application"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, database *gorm.DB, validate *validator.Validate) *gin.Engine {

	intoleranceRepo := postgres.NewIntolerancePostgres(database)

	dietRepo := postgres.NewDietPostgres(database)

	cuisineRepo := postgres.NewCuisinePostgres(database)

	routes := router.Group("/api")

	routes.POST("/intolerance", NewCreateIntolerance(
		application.NewCreateIntolerance(intoleranceRepo),
	))

	routes.GET("/intolerances", NewGetIntolerances(
		application.NewGetIntolerances(intoleranceRepo),
	))

	routes.GET("/intolerances/:intolerance_id", NewGetIntoleranceByID(
		application.NewGetIntoleranceByID(intoleranceRepo),
	))

	routes.POST("/diet", NewCreateDiet(
		application.NewCreateDiet(dietRepo),
	))

	routes.GET("/diets", NewGetDiets(
		application.NewGetDiets(dietRepo),
	))

	routes.GET("/diets/:diet_id", NewGetDietByID(
		application.NewGetDietByID(dietRepo),
	))

	routes.POST("/cuisine", NewCreateCuisine(
		application.NewCreateCuisine(cuisineRepo),
	))

	routes.GET("/cuisines", NewGetCuisines(
		application.NewGetCuisines(cuisineRepo),
	))

	routes.GET("/cuisines/:cuisine_id", NewGetCuisineByID(
		application.NewGetCuisineByID(cuisineRepo),
	))

	return router
}
