package routes

import (
	"online_store/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateProduct)
	r.GET("/", controllers.GetAllProduct)
	r.GET("/:id", controllers.GetProductById)
	r.PUT("/:id", controllers.UpdateProduct)
	r.DELETE("/:id", controllers.DeleteProduct)

}
