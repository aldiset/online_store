package routes

import (
	"online_store/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateCategory)
	r.GET("/", controllers.GetAllCategory)
	r.GET("/:id", controllers.GetCategoryById)
	r.PUT("/:id", controllers.UpdateCategory)
	r.DELETE("/:id", controllers.DeleteCategory)

}
