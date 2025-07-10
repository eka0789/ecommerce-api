package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ecommerce-api/internal/cache"
	"ecommerce-api/internal/model"
	"ecommerce-api/internal/queue"
	"ecommerce-api/internal/repository"
)

type OrderHandler struct {
	repo   repository.OrderRepository
	cache *cache.RedisClient
	log   repository.OrderLogRepository
	rmq   *queue.RabbitMQPublisher
}

func NewOrderHandler(repo repository.OrderRepository, cache *cache.RedisClient, log repository.OrderLogRepository, rmq *queue.RabbitMQPublisher) *OrderHandler {
	return &OrderHandler{repo: repo, cache: cache, log: log, rmq: rmq}
}

func (h *OrderHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders", h.GetAllOrders)
	r.GET("/orders/:id", h.GetOrderByID)
	r.DELETE("/orders/:id", h.DeleteOrder)
	r.GET("/orders/:id/logs", h.GetOrderLogs)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	orderID, err := h.repo.Create(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}
	_ = h.log.Log(orderID, "created")
	_ = h.rmq.PublishOrderID(orderID)
	c.JSON(http.StatusOK, gin.H{"order_id": orderID})
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	order, err := h.repo.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	_ = h.log.Log(id, "deleted")
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}

func (h *OrderHandler) GetOrderLogs(c *gin.Context) {
	logs, err := h.log.GetLogsByOrderID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get logs"})
		return
	}
	c.JSON(http.StatusOK, logs)
}