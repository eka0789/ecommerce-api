package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/legitdev/ecommerce-api/internal/model"
)

func TestCreateOrderHandler(t *testing.T) {
	r := gin.Default()
	mockRepo := &MockOrderRepo{}
	mockLogger := &MockLogger{}
	h := NewOrderHandler(mockRepo, nil, mockLogger, nil)
	h.RegisterRoutes(r)

	order := model.Order{UserID: "u1", Items: []model.OrderItem{{ProductID: "p1", Quantity: 1, Price: 100}}}
	body, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}

type MockOrderRepo struct{}

func (m *MockOrderRepo) Create(o *model.Order) (string, error) { return "123", nil }
func (m *MockOrderRepo) GetAll() ([]model.Order, error)        { return nil, nil }
func (m *MockOrderRepo) GetByID(id string) (*model.Order, error) { return nil, nil }
func (m *MockOrderRepo) DeleteByID(id string) error            { return nil }

type MockLogger struct{}

func (m *MockLogger) Log(orderID string, action string) error { return nil }
func (m *MockLogger) GetLogsByOrderID(orderID string) ([]model.OrderLog, error) { return nil, nil }
