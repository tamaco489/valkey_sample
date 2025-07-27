package redis

import (
	"context"
	"fmt"
	"time"
)

func (c *redisClient) RPushBatchWithPipeline(ctx context.Context, kv map[string][]string, expiration time.Duration) error {
	const batchSize = 500

	batch := make(map[string][]string)
	var count int
	for key, values := range kv {
		batch[key] = values
		count++

		if count >= batchSize {
			if err := c.executeRPushBatch(ctx, batch, expiration); err != nil {
				return fmt.Errorf("failed to execute RPUSH batch: %w", err)
			}
			batch = make(map[string][]string)
			count = 0
		}
	}

	if len(batch) > 0 {
		if err := c.executeRPushBatch(ctx, batch, expiration); err != nil {
			return fmt.Errorf("failed to execute final RPUSH batch: %w", err)
		}
	}

	return nil
}

func (c *redisClient) executeRPushBatch(ctx context.Context, batch map[string][]string, expiration time.Duration) error {
	pipe := c.rdb.Pipeline()
	for key, values := range batch {
		// RPUSHでアイテムIDをリストの右端に追加
		pipe.RPush(ctx, key, values)
		// 有効期限を設定
		pipe.Expire(ctx, key, expiration)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute RPUSH batch: %w", err)
	}
	return nil
}
