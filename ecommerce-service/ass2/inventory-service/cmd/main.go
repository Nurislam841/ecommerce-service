package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"inventory-service/config"
	grpcHandler "inventory-service/internal/adapter/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"log"
	"net"
	"proto/inventorypb"
)

func main() {
	db := config.ConnectMongo()
	defer func() {
		client := db.Client()
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Failed to disconnect from MongoDB:", err)
		}
	}()

	client := db.Client()
	productRepo := repository.NewProductRepository(client, "ecommerce")
	reviewRepo := repository.NewReviewRepository(client, "reviews")
	productUsecase := usecase.NewProductUsecase(productRepo)
	reviewUsecase := usecase.NewReviewUsecase(reviewRepo)

	grpcServer := grpc.NewServer()

	inventoryServiceServer := grpcHandler.NewInventoryServer(productUsecase, reviewUsecase)
	inventorypb.RegisterInventoryServiceServer(grpcServer, inventoryServiceServer)

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Inventory Service running on :5001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
