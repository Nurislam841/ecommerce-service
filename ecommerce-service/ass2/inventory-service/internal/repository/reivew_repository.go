package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-service/internal/entity"
	"time"
)

type ReviewRepository struct {
	collection *mongo.Collection
}

func NewReviewRepository(client *mongo.Client, dbName string) *ReviewRepository {
	return &ReviewRepository{
		collection: client.Database(dbName).Collection("reviews"),
	}
}

func (r *ReviewRepository) CreateReview(ctx context.Context, review *entity.Review) error {
	review.ID = primitive.NewObjectID().Hex()
	review.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, review)
	return err
}

func (r *ReviewRepository) GetProductReviews(ctx context.Context, productID string) ([]entity.Review, error) {
	filter := bson.M{"product_id": productID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []entity.Review
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}
