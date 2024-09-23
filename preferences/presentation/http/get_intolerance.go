package http

import (
	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/application"
	"github.com/gin-gonic/gin"
)

func NewGetIntolerances(
	getAll application.GetIntolerances,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		intolerances, err := getAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":      "intolerances fetched successfully",
			"intolerances": intolerances,
		})

	}
}
