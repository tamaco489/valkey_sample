package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (c *redisClient) SMembersBatchWithPipeline(ctx context.Context, keys []string) (map[string][]string, error) {
	const batchSize = 500
	result := make(map[string][]string)

	for i := 0; i < len(keys); i += batchSize {
		end := i + batchSize
		if end > len(keys) {
			end = len(keys)
		}

		batch := keys[i:end]
		batchResult, err := c.executeSMembersBatch(ctx, batch)
		if err != nil {
			return nil, fmt.Errorf("failed to execute SMEMBERS batch: %w", err)
		}

		for key, members := range batchResult {
			result[key] = members
		}
	}

	return result, nil
}

func (c *redisClient) executeSMembersBatch(ctx context.Context, keys []string) (map[string][]string, error) {
	pipe := c.rdb.Pipeline()

	cmds := make(map[string]*redis.StringSliceCmd)
	for _, key := range keys {
		cmds[key] = pipe.SMembers(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("failed to execute SMEMBERS batch: %w", err)
	}

	result := make(map[string][]string)
	for key, cmd := range cmds {
		members, err := cmd.Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return nil, fmt.Errorf("failed to get members for key %s: %w", key, err)
		}
		result[key] = members
	}

	return result, nil
}
