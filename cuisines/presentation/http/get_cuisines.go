package http

import (
	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/application"
	"github.com/gin-gonic/gin"
)

func NewGetCuisines(
	getAll application.GetCuisines,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cuisines, err := getAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "cuisines fetched successfully",
			"cuisines": cuisines,
		})

	}
}
