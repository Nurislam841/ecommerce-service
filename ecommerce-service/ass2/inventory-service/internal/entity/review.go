package entity

import (
	"time"
)

type Review struct {
	ID        string    `bson:"_id,omitempty"`
	ProductID string    `bson:"product_id"`
	UserID    string    `bson:"user_id"`
	Rating    int32     `bson:"rating"`
	Comment   string    `bson:"comment"`
	CreatedAt time.Time `bson:"created_at"`
}
