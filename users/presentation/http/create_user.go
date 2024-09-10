package http

import (
	"net/http"

	"github.com/Abuzar-JS/spoonacular-food-API/users/application"
	"github.com/Abuzar-JS/spoonacular-food-API/users/presentation/models"
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

		request := application.CreateUserRequest{
			Name:     body.Name,
			Location: body.Location,
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
