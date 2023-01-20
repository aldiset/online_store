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

// @title API Online Store
// @version 1.0
// @description You can visit the GitHub repository at https://github.com/aldiset/online_store

// @host localhost:80
// @BasePath /
// @query.collection.format multi
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

	r.Run(":80")
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
}
