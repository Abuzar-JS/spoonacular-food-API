package http

import (
	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/user/application"
	"github.com/gin-gonic/gin"
)

func NewSpoonacularRecipe(
	getRecipe application.SpoonacularRecipe,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cuisine := ctx.Query("cuisine")
		if cuisine == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Cuisine query parameter is required",
			})
			return
		}

		recipes, err := getRecipe(ctx, cuisine)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"recipes": recipes,
		})
	}
}
