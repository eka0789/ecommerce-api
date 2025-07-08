package repository

import (
	"context"
	"time"

	"github.com/legitdev/ecommerce-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderMongoRepo struct {
	collection *mongo.Collection
}

func NewOrderMongoRepo(db *mongo.Database) OrderRepository {
	return &orderMongoRepo{collection: db.Collection("orders")}
}

func (r *orderMongoRepo) Create(order *model.Order) (string, error) {
	order.ID = primitive.NewObjectID().Hex()
	order.CreatedAt = time.Now()
	order.Status = "pending"
	_, err := r.collection.InsertOne(context.TODO(), order)
	return order.ID, err
}

func (r *orderMongoRepo) GetAll() ([]model.Order, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var orders []model.Order
	if err := cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderMongoRepo) GetByID(id string) (*model.Order, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var order model.Order
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&order)
	return &order, err
}

func (r *orderMongoRepo) DeleteByID(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}