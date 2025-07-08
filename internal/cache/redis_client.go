package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{Addr: addr}),
	}
}

func (r *RedisClient) SetStatus(ctx context.Context, id, status string) error {
	return r.client.Set(ctx, "order:"+id, status, 10*time.Minute).Err()
}

func (r *RedisClient) GetStatus(ctx context.Context, id string) (string, error) {
	return r.client.Get(ctx, "order:"+id).Result()
}