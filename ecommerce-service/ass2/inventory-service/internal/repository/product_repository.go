package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"inventory-service/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client, dbName string) *ProductRepository {
	return &ProductRepository{
		collection: client.Database(dbName).Collection("products"),
	}
}

func (r *ProductRepository) GetProductByID(id string) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID format: %v", err)
	}

	var product entity.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, entity.ErrProductNotFound
		}
		return nil, err
	}

	return &product, nil

}

func (r *ProductRepository) CreateProduct(product *entity.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID := primitive.NewObjectID()
	product.ID = objectID.Hex()

	doc := bson.M{
		"_id":         objectID,
		"name":        product.Name,
		"price":       product.Price,
		"stock":       product.Stock,
		"category_id": product.CategoryID,
	}
	fmt.Printf("Inserting document: %+v\n", doc)
	_, err := r.collection.InsertOne(ctx, doc)
	return err
}

func (r *ProductRepository) UpdateProduct(product *entity.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"price":       product.Price,
			"stock":       product.Stock,
			"category_id": product.CategoryID,
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}

func (r *ProductRepository) DeleteProduct(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid product ID: %v", err)
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *ProductRepository) GetAllProducts(name string, category string, limit, offset int) ([]entity.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}

	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}

	if category != "" {
		filter["category_id"] = category
	}

	opts := options.Find()
	if limit > 0 {
		opts.SetLimit(int64(limit))
		opts.SetSkip(int64(offset))
	}

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []entity.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}
