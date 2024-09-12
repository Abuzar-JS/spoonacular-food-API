package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/go-spoonacular-api/diet/application"
	"github.com/gin-gonic/gin"
)

func NewGetDietByID(
	getByID application.GetDietByID,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dietID := ctx.Param("diet_id")
		ID, err := strconv.Atoi(dietID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		diet, err := getByID(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "diet found",
			"diet":    diet,
		})
	}
}
