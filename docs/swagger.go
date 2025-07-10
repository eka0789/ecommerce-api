package docs

import "github.com/swaggo/swag"

// @title Ecommerce API
// @version 1.0
// @description API for order processing using Go, MongoDB, Kafka, Redis, and RabbitMQ
// @host localhost:8080
// @BasePath /
func init() {
	swag.Register("swagger", &s{})
}

type s struct{}

func (s *s) ReadDoc() string {
	return `{
		"swagger": "2.0",
		"info": {
			"description": "API for order processing",
			"title": "Ecommerce API",
			"version": "1.0"
		},
		"host": "localhost:8080",
		"basePath": "/"
	}`
}

