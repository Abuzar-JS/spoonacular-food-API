package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/go-spoonacular-api/user_preferences/application"
	"github.com/gin-gonic/gin"
)

func NewGetUserCuisines(
	service application.GetCuisinesByUserID,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userID := ctx.Param("user_id")
		uID, err := strconv.Atoi(userID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		request := application.GetCuisinesByUserIDRequest{
			UserID: uID,
		}

		userCuisines, err := service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user_cuisines": userCuisines,
		})
	}
}
