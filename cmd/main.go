package main

import (
	"net/http"
	"online_store/internal/authentication"
	"online_store/internal/core"
	"online_store/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()

	r := gin.Default()

	home := r.Group("/")
	auth := r.Group("/api/auth")

	home.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
	})

	auth.POST("/register", core.RegisterUser)
	auth.POST("/login", core.Login)

	whoami := r.Group("/api/me")
	whoami.Use(authentication.JwtAuth())
	whoami.GET("/", core.CurrentUser)

	r.Run(":80")
}
