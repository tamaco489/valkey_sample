package redis

import (
	"context"
	"fmt"
	"time"
)

func (c *redisClient) SetWithPipeline(ctx context.Context, kv map[string]string, expiration time.Duration) error {
	pipe := c.rdb.Pipeline()
	for key, value := range kv {
		pipe.Set(ctx, key, value, expiration)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to save data to Redis: %w", err)
	}
	return nil
}
