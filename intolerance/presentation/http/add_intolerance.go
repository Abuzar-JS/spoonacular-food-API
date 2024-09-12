package http

import (
	"net/http"
	"strings"

	"github.com/Abuzar-JS/go-spoonacular-api/intolerance/application"
	"github.com/Abuzar-JS/go-spoonacular-api/intolerance/presentation/models"
	"github.com/gin-gonic/gin"
)

func NewCreateIntolerance(
	service application.CreateIntolerance,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body models.CreateIntoleranceRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		request := application.CreateIntoleranceRequest{
			Name: strings.ToLower(body.Name),
		}

		intolerance, err := service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":     "intolerance created successfully",
			"intolerance": intolerance,
		})
	}
}
