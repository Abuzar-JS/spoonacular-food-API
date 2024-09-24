package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/Abuzar-JS/go-spoonacular-api/config"
	preferenceRoutes "github.com/Abuzar-JS/go-spoonacular-api/preferences/presentation/http"
	userRoutes "github.com/Abuzar-JS/go-spoonacular-api/user/presentation/http"
	userPreferenceRoutes "github.com/Abuzar-JS/go-spoonacular-api/user_preferences/presentation/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	db := config.DatabaseConnection()
	validate := validator.New()

	ginRouter := gin.Default()

	userRoutes.RegisterRoutes(ginRouter, db, validate)

	preferenceRoutes.RegisterRoutes(ginRouter, db, validate)

	userPreferenceRoutes.RegisterRoutes(ginRouter, db, validate)

	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: ginRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server not started")
	}

}
