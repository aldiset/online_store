package routes

import (
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	r.POST("/")
	r.GET("/")
	r.GET("/:id")
	r.PUT("/:id")
	r.DELETE("/:id")

}
