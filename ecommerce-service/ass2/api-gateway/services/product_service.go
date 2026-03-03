package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"proto/inventorypb"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

var inventoryClient inventorypb.InventoryServiceClient

func init() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to inventory service: %v", err)
	}
	inventoryClient = inventorypb.NewInventoryServiceClient(conn)
}

func CreateProduct(c *gin.Context) {
	var input inventorypb.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	fmt.Println(input)
	res, err := inventoryClient.CreateProduct(context.Background(), &inventorypb.CreateProductRequest{Product: &input})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": res.Id})
}

func GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	res, err := inventoryClient.GetProduct(context.Background(), &inventorypb.GetProductRequest{Id: int32(id)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, res.Product)
}

func UpdateProduct(c *gin.Context) {
	var input inventorypb.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	res, err := inventoryClient.UpdateProduct(context.Background(), &inventorypb.UpdateProductRequest{Product: &input})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}
	c.JSON(http.StatusOK, res.Product)
}

func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	res, err := inventoryClient.DeleteProduct(context.Background(), &inventorypb.DeleteProductRequest{Id: string(id)})
	if err != nil || !res.Success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
func GetProducts(c *gin.Context) {
	name := c.Query("name")
	category, _ := strconv.Atoi(c.DefaultQuery("category", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	req := &inventorypb.ListProductsRequest{
		Name:     name,
		Category: int32(category),
		Limit:    int32(limit),
		Offset:   int32(offset),
	}

	res, err := inventoryClient.ListProducts(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, res.Products)
}
