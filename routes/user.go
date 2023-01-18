package routes

import (
	"online_store/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllUser)
	r.GET("/:id", controllers.GetUserById)
	r.PUT("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)

}
