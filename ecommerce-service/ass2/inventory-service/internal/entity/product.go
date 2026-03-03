package entity

type Product struct {
	ID         string  `bson:"_id,omitempty" json:"id"`
	Name       string  `bson:"name" json:"name"`
	Price      float64 `bson:"price" json:"price"`
	Stock      int32   `bson:"stock" json:"stock"`
	CategoryID string  `bson:"category_id" json:"category_id"`
}
