package routes

import (
	"online_store/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateCart)
	r.GET("/", controllers.GetAllCart)
	r.GET("/:id", controllers.GetCartById)
	r.PUT("/:id", controllers.UpdateCart)
	r.DELETE("/:id", controllers.DeleteCart)

}
