package http

import (
	"net/http"
	"strings"

	"github.com/Abuzar-JS/go-spoonacular-api/users/application"
	"github.com/Abuzar-JS/go-spoonacular-api/users/presentation/models"
	"github.com/gin-gonic/gin"
)

func NewCreateUser(
	service application.CreateUser,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body models.CreateUserRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		nameLower := strings.ToLower(body.Name)

		request := application.CreateUserRequest{
			Name:     nameLower,
			Password: body.Password,
		}

		user, err := service(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, returnError(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "user created successfully",
			"user":    user,
		})
	}
}
