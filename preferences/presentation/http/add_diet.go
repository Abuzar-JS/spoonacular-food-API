package http

import (
	"net/http"
	"strings"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/application"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/presentation/models"
	"github.com/gin-gonic/gin"
)

func NewCreateDiet(
	service application.CreateDiet,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body models.CreateDietRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		request := application.CreateDietRequest{
			Name: strings.ToLower(body.Name),
		}

		diet, err := service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "diet created successfully",
			"diet":    diet,
		})
	}
}
