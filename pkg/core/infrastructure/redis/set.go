package redis

import (
	"context"
	"time"
)

func (c *redisClient) SetWithPipeline(ctx context.Context, kv map[string]string, expiration time.Duration) error {
	pipe := c.rdb.Pipeline()
	for key, value := range kv {
		pipe.Set(ctx, key, value, expiration)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}
	return nil
}
