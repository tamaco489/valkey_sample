package redis

import (
	"context"
	"fmt"
	"time"
)

func (c *redisClient) SAddBatchWithPipeline(ctx context.Context, kv map[string][]string, expiration time.Duration) error {
	const batchSize = 500

	batch := make(map[string][]string)
	var count int
	for key, values := range kv {
		batch[key] = values
		count++

		if count >= batchSize {
			if err := c.executeSAddBatch(ctx, batch, expiration); err != nil {
				return fmt.Errorf("failed to execute SADD batch: %w", err)
			}
			batch = make(map[string][]string)
			count = 0
		}
	}

	if len(batch) > 0 {
		if err := c.executeSAddBatch(ctx, batch, expiration); err != nil {
			return fmt.Errorf("failed to execute final SADD batch: %w", err)
		}
	}

	return nil
}

func (c *redisClient) executeSAddBatch(ctx context.Context, batch map[string][]string, expiration time.Duration) error {
	pipe := c.rdb.Pipeline()
	for key, values := range batch {
		pipe.SAdd(ctx, key, values)
		pipe.Expire(ctx, key, expiration)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute SADD batch: %w", err)
	}
	return nil
}
