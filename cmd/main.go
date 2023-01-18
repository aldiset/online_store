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
	whoami := r.Group("/api/me")
	product := r.Group("/api/product")
	category := r.Group("/api/category")
	cart := r.Group("/api/cart")
	transaction := r.Group("/api/transaction")

	home.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
	})

	auth.POST("/register", core.RegisterUser)
	auth.POST("/login", core.Login)

	whoami.Use(authentication.JwtAuth())
	whoami.GET("/", core.CurrentUser)

	product.Use(authentication.JwtAuth())
	product.GET("/")
	product.GET("/:id")
	product.POST("/:id")
	product.PUT("/:id")
	product.DELETE("/:id")

	category.Use(authentication.JwtAuth())
	category.GET("/")
	category.GET("/:id")
	category.POST("/:id")
	category.PUT("/:id")
	category.DELETE("/:id")

	cart.Use(authentication.JwtAuth())
	cart.GET("/")
	cart.GET("/:id")
	cart.POST("/:id")
	cart.PUT("/:id")
	cart.DELETE("/:id")

	transaction.Use(authentication.JwtAuth())
	transaction.GET("/")
	transaction.GET("/:id")
	transaction.POST("/:id")
	transaction.PUT("/:id")
	transaction.DELETE("/:id")

	r.Run(":80")
}
