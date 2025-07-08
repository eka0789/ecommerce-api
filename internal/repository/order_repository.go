package repository

import "github.com/legitdev/ecommerce-api/internal/model"

type OrderRepository interface {
	Create(*model.Order) (string, error)
	GetAll() ([]model.Order, error)
	GetByID(id string) (*model.Order, error)
	DeleteByID(id string) error
}

type OrderLogRepository interface {
	Log(orderID, action string) error
	GetLogsByOrderID(orderID string) ([]model.OrderLog, error)
}