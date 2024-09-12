package http

import (
	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/users/application"
	"github.com/gin-gonic/gin"
)

func NewGetUsers(
	getAll application.GetUsers,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := getAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "users fetched successfully",
			"users":   users,
		})

	}
}
