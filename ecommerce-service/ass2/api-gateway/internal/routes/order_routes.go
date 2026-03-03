package routes

import (
	"github.com/AskatNa/SecondAssignment/services"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine) {
	r.POST("/orders", services.CreateOrder)
	r.GET("/orders/:id", services.GetOrderByID)
	r.PATCH("/orders/:id", services.UpdateOrderStatus)
	r.GET("/orders", services.GetUserOrders)
}
