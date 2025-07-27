package redis

import (
	"context"
	"time"
)

// RedisService represents a Redis service interface
type RedisService interface {
	Ping(ctx context.Context) error
	SetBatchWithPipeline(ctx context.Context, kv map[string]string, expiration time.Duration) error
	SAddBatchWithPipeline(ctx context.Context, kv map[string][]string, expiration time.Duration) error
	Close() error
}

var _ RedisService = (*redisClient)(nil)
