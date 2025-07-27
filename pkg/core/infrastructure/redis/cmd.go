package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Get retrieves a value from Redis
func (c *client) Get(ctx context.Context, key string) (string, error) {
	result, err := c.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return result, err
}

// Set stores a value in Redis
func (c *client) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return c.rdb.Set(ctx, key, value, expiration).Err()
}

// Del removes keys from Redis
func (c *client) Del(ctx context.Context, keys ...string) error {
	return c.rdb.Del(ctx, keys...).Err()
}

// Exists checks if keys exist in Redis
func (c *client) Exists(ctx context.Context, keys ...string) (int64, error) {
	return c.rdb.Exists(ctx, keys...).Result()
}

// Ping tests the Redis connection
func (c *client) Ping(ctx context.Context) error {
	return c.rdb.Ping(ctx).Err()
}

// Close closes the Redis connection
func (c *client) Close() error {
	return c.rdb.Close()
}
