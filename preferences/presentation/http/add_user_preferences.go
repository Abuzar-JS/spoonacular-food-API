package http

import (
	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/application"
	"github.com/Abuzar-JS/go-spoonacular-api/preferences/presentation/models"
	"github.com/gin-gonic/gin"
)

func NewAddUserPreferences(
	service application.AddUserPreferences,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body models.AddUserPreferencesRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		request := application.AddUserPreferencesRequest{
			UserID:       body.UserID,
			Cuisines:     body.CuisineID,
			Diets:        body.DietID,
			Intolerances: body.IntoleranceID,
		}

		preferences, err := service(ctx, request)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":     "preferences saved successfully",
			"preferences": preferences,
		})
	}
}
