package main

import (
	"net/http"
	"online_store/middleware"
	"online_store/models"
	"online_store/routes"

	docs "online_store/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	models.ConnectDB()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	index := r.Group("/")
	api := r.Group("/api")
	index.GET("/", Index)
	routes.AuthenticationRoutes(api.Group("/auth"))

	api.Use(middleware.JwtAuth())
	routes.UserRoutes(api.Group("/user"))
	routes.CategoryRoutes(api.Group("/category"))
	routes.PaymentRoutes(api.Group("/payment"))
	routes.ProductRoutes(api.Group("/product"))
	routes.CartRoutes(api.Group("/cart"))
	routes.TransactionRoutes(api.Group("/transaction"))

	r.Run("0.0.0.0:8080")
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
}
