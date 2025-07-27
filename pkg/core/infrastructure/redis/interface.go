package redis

import (
	"context"
	"time"
)

// RedisService represents a Redis service interface
type RedisService interface {
	Ping(ctx context.Context) error
	SetBatchWithPipeline(ctx context.Context, kv map[string]string, expiration time.Duration) error
	GetBatchWithPipeline(ctx context.Context, keys []string) (map[string]string, error)
	SAddBatchWithPipeline(ctx context.Context, kv map[string][]string, expiration time.Duration) error
	SMembersBatchWithPipeline(ctx context.Context, keys []string) (map[string][]string, error)
	RPushBatchWithPipeline(ctx context.Context, kv map[string][]string, expiration time.Duration) error
	LRangeBatchWithPipeline(ctx context.Context, keys []string) (map[string][]string, error)
	Close() error
}

var _ RedisService = (*redisClient)(nil)
