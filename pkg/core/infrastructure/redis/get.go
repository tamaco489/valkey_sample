package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (c *redisClient) GetBatchWithPipeline(ctx context.Context, keys []string) (map[string]string, error) {
	const batchSize = 500
	result := make(map[string]string)

	for i := 0; i < len(keys); i += batchSize {
		end := i + batchSize
		if end > len(keys) {
			end = len(keys)
		}

		batch := keys[i:end]
		batchResult, err := c.executeGetBatch(ctx, batch)
		if err != nil {
			return nil, fmt.Errorf("failed to execute GET batch: %w", err)
		}

		for key, value := range batchResult {
			result[key] = value
		}
	}

	return result, nil
}

func (c *redisClient) executeGetBatch(ctx context.Context, keys []string) (map[string]string, error) {
	pipe := c.rdb.Pipeline()

	cmds := make(map[string]*redis.StringCmd)
	for _, key := range keys {
		cmds[key] = pipe.Get(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("failed to execute GET batch: %w", err)
	}

	result := make(map[string]string)
	for key, cmd := range cmds {
		val, err := cmd.Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return nil, fmt.Errorf("failed to get value for key %s: %w", key, err)
		}
		result[key] = val
	}

	return result, nil
}
