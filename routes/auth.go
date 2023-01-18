package routes

import (
	"online_store/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(r *gin.RouterGroup) {
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.Login)
	r.GET("/me", controllers.CurrentUser)

}
