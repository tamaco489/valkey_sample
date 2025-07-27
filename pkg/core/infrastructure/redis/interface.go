package redis

import (
	"context"
	"time"
)

// Client represents a Redis client interface
type RedisService interface {
	Ping(ctx context.Context) error
	SetWithPipeline(ctx context.Context, kv map[string]string, expiration time.Duration) error
	Close() error
}

var _ RedisService = (*redisClient)(nil)
