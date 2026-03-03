package services

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"proto/orderpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var orderClient orderpb.OrderServiceClient

func init() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	orderClient = orderpb.NewOrderServiceClient(conn)
}
func CreateOrder(c *gin.Context) {
	var input orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	resp, err := orderClient.CreateOrder(context.Background(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order_id": resp.Order.Id})
}

func GetOrderByID(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))
	resp, err := orderClient.GetOrder(context.Background(), &orderpb.GetOrderRequest{Id: int32(orderID)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order_id": resp.Order.Id,
		"user_id":  resp.Order.UserId,
		"status":   resp.Order.Status,
	})
}

func UpdateOrderStatus(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))
	var input orderpb.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	input.Id = int32(orderID)
	_, err := orderClient.UpdateOrderStatus(context.Background(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order status updated"})
}

func GetUserOrders(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	resp, err := orderClient.GetOrdersByUser(context.Background(), &orderpb.GetOrdersByUserRequest{UserId: int32(userID)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": resp.Orders})
}
