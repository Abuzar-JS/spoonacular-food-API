package http

import (
	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/diet/application"
	"github.com/gin-gonic/gin"
)

func NewGetDiets(
	getAll application.GetDiets,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		diets, err := getAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "diets fetched successfully",
			"diets":   diets,
		})

	}
}
