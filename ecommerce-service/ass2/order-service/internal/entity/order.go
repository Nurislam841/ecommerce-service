package entity

type Order struct {
	ID string `bson:"_id,omitempty"`
	//UserID string      `bson:"user_id"`
	Status string      `bson:"status"`
	Items  []OrderItem `bson:"items"`
}

type OrderItem struct {
	ProductID string `bson:"product_id"`
	Quantity  int32  `bson:"quantity"`
}
