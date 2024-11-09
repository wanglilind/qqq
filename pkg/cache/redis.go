package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(config *Config) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         config.RedisAddr,
		Password:     config.RedisPassword,
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		MaxRetries:   3,
	})

	// 测试连接
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{client: client}, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
} 
