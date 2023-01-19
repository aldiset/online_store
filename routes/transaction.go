package routes

import (
	"online_store/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateTransaction)
	r.GET("/", controllers.GetAllTransaction)
	r.GET("/:id", controllers.GetTransactionById)
	r.PUT("/:id", controllers.UpdateTransaction)
	r.DELETE("/:id", controllers.DeleteTransaction)

}
