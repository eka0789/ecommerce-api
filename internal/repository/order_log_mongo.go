package repository

import (
	"context"
	"time"

	"github.com/legitdev/ecommerce-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderLogRepo struct {
	collection *mongo.Collection
}

func NewOrderLogRepo(db *mongo.Database) OrderLogRepository {
	return &orderLogRepo{collection: db.Collection("order_logs")}
}

func (r *orderLogRepo) Log(orderID string, action string) error {
	entry := model.OrderLog{
		OrderID:   orderID,
		Action:    action,
		Timestamp: time.Now(),
	}
	_, err := r.collection.InsertOne(context.TODO(), entry)
	return err
}

func (r *orderLogRepo) GetLogsByOrderID(orderID string) ([]model.OrderLog, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{"order_id": orderID})
	if err != nil {
		return nil, err
	}
	var logs []model.OrderLog
	err = cursor.All(context.TODO(), &logs)
	return logs, err
}
