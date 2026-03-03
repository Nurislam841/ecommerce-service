package usecase

import (
	"context"
	"inventory-service/internal/entity"
	"inventory-service/internal/repository"
)

type ReviewUsecase struct {
	repo *repository.ReviewRepository
}

func NewReviewUsecase(repo *repository.ReviewRepository) *ReviewUsecase {
	return &ReviewUsecase{repo: repo}
}

func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *entity.Review) error {
	return uc.repo.CreateReview(ctx, review)
}

func (uc *ReviewUsecase) GetProductReviews(ctx context.Context, productID string) ([]entity.Review, error) {
	return uc.repo.GetProductReviews(ctx, productID)
}
