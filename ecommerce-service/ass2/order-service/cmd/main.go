package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"order-service/config"
	grpcHandler "order-service/internal/adapter/grpc"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"proto/orderpb"
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
	repo := repository.NewOrderRepository(client, "ecommerce")
	orderUsecase := usecase.NewOrderUsecase(repo)

	grpcServer := grpc.NewServer()
	orderServiceServer := grpcHandler.NewOrderServer(orderUsecase)
	orderpb.RegisterOrderServiceServer(grpcServer, orderServiceServer)

	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Order Service running on :5002")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
