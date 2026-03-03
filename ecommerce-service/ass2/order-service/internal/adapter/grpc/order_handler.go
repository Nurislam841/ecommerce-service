package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"order-service/internal/usecase"
	"proto/orderpb"
)

type OrderServer struct {
	orderpb.UnimplementedOrderServiceServer
	Usecase *usecase.OrderUsecase
}

func NewOrderServer(usecase *usecase.OrderUsecase) *OrderServer {
	return &OrderServer{Usecase: usecase}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	order, err := s.Usecase.CreateOrderFromRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %v", err)
	}

	return &orderpb.CreateOrderResponse{Id: order.ID}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	order, err := s.Usecase.GetOrderByIDPB(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
	}

	return &orderpb.GetOrderResponse{Order: order}, nil
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.UpdateOrderStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	err := s.Usecase.UpdateOrderStatus(req.Id, req.Status)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update order status: %v", err)
	}

	order, err := s.Usecase.GetOrderByIDPB(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch updated order: %v", err)
	}

	return &orderpb.UpdateOrderStatusResponse{
		Success: true,
		Order:   order,
	}, nil
}
