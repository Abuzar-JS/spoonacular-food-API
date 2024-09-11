package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegisterRoutes(router *gin.Engine, validate *validator.Validate) {

	spoonacularRouter := router.Group("https://api.spoonacular.com")
	spoonacularRouter.GET("/recipes/complexSearch/:apiKey",)



}
