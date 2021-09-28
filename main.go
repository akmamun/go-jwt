package main

import (
	"net/http"
	"os"

	middleware "go-jwt/src/middleware"
	routes "go-jwt/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": nil, "message": "successfully login"})

	})

	router.Run(":" + port)
}
