package http

import (
	"net/http"
	"strconv"

	"github.com/Abuzar-JS/go-spoonacular-api/preferences/application"
	"github.com/gin-gonic/gin"
)

func NewGetIntoleranceByID(
	getByID application.GetIntoleranceByID,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		intoleranceID := ctx.Param("intolerance_id")
		ID, err := strconv.Atoi(intoleranceID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		intolerance, err := getByID(ctx.Request.Context(), ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":     "intolerance found",
			"intolerance": intolerance,
		})
	}
}
