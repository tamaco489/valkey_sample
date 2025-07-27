package redis

import (
	"context"
	"fmt"
	"time"
)

func (c *redisClient) SetBatchWithPipeline(ctx context.Context, kv map[string]string, expiration time.Duration) error {
	const batchSize = 500

	batch := make(map[string]string)
	var count int
	for key, value := range kv {
		batch[key] = value
		count++

		if count >= batchSize {
			if err := c.executeBatch(ctx, batch, expiration); err != nil {
				return fmt.Errorf("failed to execute batch: %w", err)
			}
			batch = make(map[string]string)
			count = 0
		}
	}

	if len(batch) > 0 {
		if err := c.executeBatch(ctx, batch, expiration); err != nil {
			return fmt.Errorf("failed to execute final batch: %w", err)
		}
	}

	return nil
}

func (c *redisClient) executeBatch(ctx context.Context, batch map[string]string, expiration time.Duration) error {
	pipe := c.rdb.Pipeline()
	for key, value := range batch {
		pipe.Set(ctx, key, value, expiration)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute batch: %w", err)
	}
	return nil
}
