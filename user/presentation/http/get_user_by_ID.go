package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/go-spoonacular-api/user/application"
	"github.com/gin-gonic/gin"
)

func NewGetUserByID(
	getByID application.GetUserByID,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")
		ID, err := strconv.Atoi(userID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := getByID(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user found",
			"user":    user,
		})
	}
}
