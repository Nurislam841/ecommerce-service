package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"inventory-service/internal/entity"
	"inventory-service/internal/usecase"
	"proto/inventorypb"
	"time"
)

type InventoryServer struct {
	inventorypb.UnimplementedInventoryServiceServer
	ProductUsecase *usecase.ProductUsecase
	ReviewUsecase  *usecase.ReviewUsecase
}

func NewInventoryServer(productUC *usecase.ProductUsecase, reviewUC *usecase.ReviewUsecase) *InventoryServer {
	return &InventoryServer{
		ProductUsecase: productUC,
		ReviewUsecase:  reviewUC,
	}
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *inventorypb.CreateProductRequest) (*inventorypb.CreateProductResponse, error) {
	if req == nil || req.Product == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}
	fmt.Printf("Received product: Name=%s, Price=%f, Stock=%d, Category=%s\n",
		req.Product.Name, req.Product.Price, req.Product.Stock, req.Product.CategoryId)

	p := req.Product
	product := entity.Product{
		Name:       p.Name,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryId,
	}

	fmt.Println("Creating product:", product)

	createdProduct, err := s.ProductUsecase.CreateProduct(product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %v", err)
	}

	return &inventorypb.CreateProductResponse{Id: createdProduct.ID}, nil
}

func (s *InventoryServer) GetProduct(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.GetProductResponse, error) {
	fmt.Printf("GetProduct request received for ID: %s\n", req.Id)
	product, err := s.ProductUsecase.GetProduct(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	return &inventorypb.GetProductResponse{
		Product: &inventorypb.Product{
			Id:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryId: product.CategoryID,
		},
	}, nil
}

func (s *InventoryServer) ListProducts(ctx context.Context, req *inventorypb.ListProductsRequest) (*inventorypb.ListProductsResponse, error) {
	products, err := s.ProductUsecase.GetAllProducts(req.Name, req.Category, int(req.Limit), int(req.Offset)) // Category now string
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %v", err)
	}

	var pbProducts []*inventorypb.Product
	for _, product := range products {
		pbProducts = append(pbProducts, &inventorypb.Product{
			Id:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryId: product.CategoryID,
		})
	}
	resp := &inventorypb.ListProductsResponse{Products: pbProducts}
	fmt.Printf("Sending ListProductsResponse with %d products\n", len(pbProducts))
	return resp, nil
}

func (s *InventoryServer) UpdateProduct(ctx context.Context, req *inventorypb.UpdateProductRequest) (*inventorypb.UpdateProductResponse, error) {
	if req == nil || req.Product == nil {
		return nil, status.Error(codes.InvalidArgument, "request or product cannot be nil")
	}

	product := &entity.Product{
		ID:         req.Product.Id,
		Name:       req.Product.Name,
		Price:      req.Product.Price,
		Stock:      req.Product.Stock,
		CategoryID: req.Product.CategoryId,
	}

	err := s.ProductUsecase.UpdateProduct(product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}
	return &inventorypb.UpdateProductResponse{
		Success: true,
		Product: &inventorypb.Product{
			Id:         product.ID,
			Name:       product.Name,
			Price:      float64(product.Price),
			Stock:      product.Stock,
			CategoryId: product.CategoryID,
		},
	}, nil
}

func (s *InventoryServer) DeleteProduct(ctx context.Context, req *inventorypb.DeleteProductRequest) (*inventorypb.DeleteProductResponse, error) {
	err := s.ProductUsecase.DeleteProduct(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete product: %v", err)
	}
	return &inventorypb.DeleteProductResponse{Success: true}, nil
}

func (s *InventoryServer) CreateReview(ctx context.Context, req *inventorypb.CreateReviewRequest) (*inventorypb.CreateReviewResponse, error) {
	review := &entity.Review{
		ProductID: req.ProductId,
		UserID:    req.UserId,
		Rating:    req.Rating,
		Comment:   req.Comment,
		CreatedAt: time.Now(),
	}
	if err := s.ReviewUsecase.CreateReview(ctx, review); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create review: %v", err)
	}
	return &inventorypb.CreateReviewResponse{
		Review: &inventorypb.Review{
			Id:        review.ID,
			ProductId: review.ProductID,
			UserId:    review.UserID,
			Rating:    review.Rating,
			Comment:   review.Comment,
			CreatedAt: review.CreatedAt.Format(time.RFC1123),
		},
	}, nil
}
func (s *InventoryServer) GetProductReview(ctx context.Context, req *inventorypb.GetProductReviewRequest) (*inventorypb.GetProductReviewResponse, error) {
	reviews, err := s.ReviewUsecase.GetProductReviews(ctx, req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get reviews: %v", err)
	}
	pbReviews := make([]*inventorypb.Review, len(reviews))
	for i, review := range reviews {
		pbReviews[i] = &inventorypb.Review{
			Id:        review.ID,
			ProductId: review.ProductID,
			UserId:    review.UserID,
			Rating:    review.Rating,
			Comment:   review.Comment,
			CreatedAt: review.CreatedAt.Format(time.RFC1123),
		}
	}
	return &inventorypb.GetProductReviewResponse{Reviews: pbReviews}, nil
}
