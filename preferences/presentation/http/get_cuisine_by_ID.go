package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/application"
	"github.com/gin-gonic/gin"
)

func NewGetCuisineByID(
	getByID application.GetCuisineByID,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cuisineID := ctx.Param("cuisine_id")
		ID, err := strconv.Atoi(cuisineID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		cuisine, err := getByID(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "cuisine found",
			"cuisine": cuisine,
		})
	}
}
