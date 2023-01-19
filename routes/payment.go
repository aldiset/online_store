package routes

import (
	"online_store/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CreatePayment)
	r.GET("/", controllers.GetAllPayment)
	r.GET("/:id", controllers.GetPaymentById)
	r.PUT("/:id", controllers.UpdatePayment)
	r.DELETE("/:id", controllers.DeletePayment)

}
