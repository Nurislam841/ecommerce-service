package main

import (
	"github.com/AskatNa/SecondAssignment/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("API Gateway running on :8080")
	r.Run(":8080")
}
