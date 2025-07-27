package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (c *redisClient) LRangeBatchWithPipeline(ctx context.Context, keys []string) (map[string][]string, error) {
	const batchSize = 500
	result := make(map[string][]string)

	for i := 0; i < len(keys); i += batchSize {
		end := i + batchSize
		if end > len(keys) {
			end = len(keys)
		}

		batch := keys[i:end]
		batchResult, err := c.executeLRangeBatch(ctx, batch)
		if err != nil {
			return nil, fmt.Errorf("failed to execute LRANGE batch: %w", err)
		}

		for key, elements := range batchResult {
			result[key] = elements
		}
	}

	return result, nil
}

func (c *redisClient) executeLRangeBatch(ctx context.Context, keys []string) (map[string][]string, error) {
	pipe := c.rdb.Pipeline()

	cmds := make(map[string]*redis.StringSliceCmd)
	for _, key := range keys {
		// LRANGE key 0 -1 でリストの全要素を取得
		cmds[key] = pipe.LRange(ctx, key, 0, -1)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("failed to execute LRANGE batch: %w", err)
	}

	result := make(map[string][]string)
	for key, cmd := range cmds {
		elements, err := cmd.Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return nil, fmt.Errorf("failed to get list elements for key %s: %w", key, err)
		}
		result[key] = elements
	}

	return result, nil
}
