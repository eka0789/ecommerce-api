package main

import (
	"log"
	"time"

	"ecommerce-api/config"
	"ecommerce-api/internal/model"
	"ecommerce-api/internal/repository"
)

func main() {
	log.Println("🚀 Starting ecommerce seeder...")

	cfg := config.LoadConfig()
	mongoClient := config.ConnectMongo(cfg)
	if mongoClient == nil {
		log.Fatal("❌ Failed to connect to MongoDB")
	}
	log.Println("✅ MongoDB connected")

	orderRepo := repository.NewOrderMongoRepo(mongoClient)

	orders := []model.Order{
		{
			UserID: "u001",
			Items: []model.OrderItem{
				{ProductID: "PRD001", Quantity: 1, Price: 749000}, // Logitech MX Master 3
				{ProductID: "PRD002", Quantity: 2, Price: 129000}, // Sandisk Flashdisk 64GB
			},
			Total:     749000 + 2*129000,
			Status:    "pending",
			CreatedAt: time.Now(),
		},
		{
			UserID: "u002",
			Items: []model.OrderItem{
				{ProductID: "PRD003", Quantity: 1, Price: 1350000}, // Redmi Note 12
			},
			Total:     1350000,
			Status:    "confirmed",
			CreatedAt: time.Now(),
		},
		{
			UserID: "u003",
			Items: []model.OrderItem{
				{ProductID: "PRD004", Quantity: 1, Price: 549000}, // Philips Air Fryer
				{ProductID: "PRD005", Quantity: 3, Price: 89000},  // Tupperware Set
			},
			Total:     549000 + 3*89000,
			Status:    "shipped",
			CreatedAt: time.Now(),
		},
		{
			UserID: "u004",
			Items: []model.OrderItem{
				{ProductID: "PRD006", Quantity: 1, Price: 999000}, // ASUS Mechanical Keyboard
			},
			Total:     999000,
			Status:    "cancelled",
			CreatedAt: time.Now(),
		},
		{
			UserID: "u005",
			Items: []model.OrderItem{
				{ProductID: "PRD007", Quantity: 2, Price: 349000}, // JBL GO 3 Speaker
			},
			Total:     2 * 349000,
			Status:    "pending",
			CreatedAt: time.Now(),
		},
	}

	for _, order := range orders {
		id, err := orderRepo.Create(&order)
		if err != nil {
			log.Printf("❌ Gagal insert order untuk user %s: %v\n", order.UserID, err)
		} else {
			log.Printf("✅ Order berhasil untuk user %s, ID: %s\n", order.UserID, id)
		}
	}

	log.Println("🎉 Seeder selesai.")
}
