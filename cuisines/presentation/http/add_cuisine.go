package http

import (
	"net/http"
	"strings"

	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/application"
	"github.com/Abuzar-JS/go-spoonacular-api/cuisines/presentation/models"
	"github.com/gin-gonic/gin"
)

func NewCreateCuisine(
	service application.CreateCuisine,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body models.CreateCuisineRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		request := application.CreateCuisineRequest{
			Name: strings.ToLower(body.Name),
		}

		cuisine, err := service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "cuisine created successfully",
			"cuisine": cuisine,
		})
	}
}
