package http

import (
	"net/http"
	"order-service/internal/entity"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	uc *usecase.OrderUsecase
}

func NewOrderHandler(r *gin.Engine, uc *usecase.OrderUsecase) {
	h := &OrderHandler{uc: uc}
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders/:id", h.GetOrder)
	r.PATCH("/orders/:id/status", h.UpdateStatus)
}

type createRequest struct {
	UserID string             `json:"user_id"`
	Items  []entity.OrderItem `json:"items"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := &entity.Order{
		//UserID: req.UserID,
		Status: "pending",
		Items:  req.Items,
	}

	if err := h.uc.CreateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	order, err := h.uc.GetOrderByIDPB(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

type statusRequest struct {
	Status string `json:"status"`
}

func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	var req statusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.UpdateOrderStatus(c.Param("id"), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
