package http

import "github.com/gin-gonic/gin"

func returnError(err error) any {
	return gin.H{
		"message": err.Error(),
	}
}
