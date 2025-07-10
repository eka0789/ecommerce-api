// Swagger annotations example in handler (CreateOrder)
// @Summary Create new order
// @Description Create a new order and publish to RabbitMQ
// @Tags orders
// @Accept json
// @Produce json
// @Param order body model.Order true "Order Payload"
// @Success 200 {object} map[string]string
// @Router /orders [post]

// Folder: internal/repository/order_mongo_test.go
package repository

import (
	"context"
	"testing"
	"time"

	"ecommerce-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateOrder(t *testing.T) {
	ml := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer ml.Close()

	db := ml.DB
	repo := NewOrderMongoRepo(db)
	order := &model.Order{
		UserID: "user123",
		Items:  []model.OrderItem{{ProductID: "p1", Quantity: 2, Price: 100}},
		Total:  200,
		Status: "pending",
	}
	id, err := repo.Create(order)
	if err != nil || id == "" {
		t.Errorf("expected successful creation, got err=%v", err)
	}
}