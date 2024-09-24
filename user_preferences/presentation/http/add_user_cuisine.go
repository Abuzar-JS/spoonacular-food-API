package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/application"
	"github.com/gin-gonic/gin"
)

func NewAddUserCuisine(
	service application.CreateUserCuisine,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cuisineID := ctx.Param("cuisine_id")
		cID, err := strconv.Atoi(cuisineID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		userID := ctx.Param("user_id")
		uID, err := strconv.Atoi(userID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		request := application.CreateUserCuisineRequest{
			UserID:    uID,
			CuisineID: cID,
		}

		_, err = service(ctx.Request.Context(), request)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "cuisine added to user successfully",
		})
	}
}
