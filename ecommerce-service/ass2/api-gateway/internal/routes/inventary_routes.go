package routes

import (
	"github.com/AskatNa/SecondAssignment/services"
	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(r *gin.Engine) {
	r.GET("/products", services.GetProducts)
	r.GET("/products/:id", services.GetProductByID)
	r.POST("/products", services.CreateProduct)
	r.PATCH("/products/:id", services.UpdateProduct)
	r.DELETE("/products/:id", services.DeleteProduct)
}
