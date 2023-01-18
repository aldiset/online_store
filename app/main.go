package main

import (
	"net/http"
	"online_store/middleware"
	"online_store/models"
	"online_store/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()

	r := gin.Default()

	index := r.Group("/")
	api := r.Group("/api")

	routes.AuthenticationRoutes(api.Group("/auth"))

	api.Use(middleware.JwtAuth())
	routes.UserRoutes(api.Group("/user"))
	routes.CategoryRoutes(api.Group("/category"))
	routes.PaymentRoutes(api.Group("/payment"))
	routes.ProductRoutes(api.Group("/product"))
	routes.TransactionRoutes(api.Group("/transaction"))

	index.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
	})

	r.Run(":80")
}
