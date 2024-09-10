package main

import (
	"fmt"

	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/config"
	userRoutes "github.com/Abuzar-JS/go-spoonacular-api/users/presentation/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	db := config.DatabaseConnection()
	validate := validator.New()

	ginRouter := gin.Default()

	// http.RegisterRoutes(ginRouter, db, validate)
	userRoutes.RegisterRoutes(ginRouter, db, validate)

	server := &http.Server{
		Addr:    ":8000",
		Handler: ginRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server not started")
	}

}
