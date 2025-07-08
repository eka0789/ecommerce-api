package model

import "time"

type OrderItem struct {
	ProductID string  `bson:"product_id" json:"product_id"`
	Quantity  int     `bson:"quantity" json:"quantity"`
	Price     float64 `bson:"price" json:"price"`
}

type Order struct {
	ID        string      `bson:"_id,omitempty" json:"id"`
	UserID    string      `bson:"user_id" json:"user_id"`
	Items     []OrderItem `bson:"items" json:"items"`
	Total     float64     `bson:"total" json:"total"`
	Status    string      `bson:"status" json:"status"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
}

type OrderLog struct {
	ID        string    `bson:"_id,omitempty"`
	OrderID   string    `bson:"order_id"`
	Action    string    `bson:"action"`
	Timestamp time.Time `bson:"timestamp"`
}