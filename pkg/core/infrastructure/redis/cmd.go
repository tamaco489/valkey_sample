package redis

import (
	"context"
)

// Ping tests the Redis connection
func (c *redisClient) Ping(ctx context.Context) error {
	return c.rdb.Ping(ctx).Err()
}

// Close closes the Redis connection
func (c *redisClient) Close() error {
	return c.rdb.Close()
}
