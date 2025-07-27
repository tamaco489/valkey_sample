package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/tamaco489/valkey_sample/pkg/core/infrastructure/redis"
)

// batchProcessor handles batch operations
type batchProcessor struct {
	userItemMap map[string]string // キー: customer:{customer_id}:user:{user_id}:items, 値: item_idのカンマ区切り
	redisClient redis.RedisService
}

// newBatchProcessor creates a new batch processor
func newBatchProcessor() *batchProcessor {
	cfg := redis.LoadConfig()
	redisAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	redisClient, err := redis.NewClient(redisAddr, cfg.Password, cfg.DB)
	if err != nil {
		slog.Error("failed to create Redis client", "error", err)
		os.Exit(1)
	}

	return &batchProcessor{
		userItemMap: make(map[string]string),
		redisClient: redisClient,
	}
}

// formatRedisKey creates a Redis key with the specified format
func (bp *batchProcessor) formatRedisKey(customerID, userID uint32) string {
	return fmt.Sprintf(RedisKeyFormat, customerID, userID)
}

// processBatch processes the entire batch
func (bp *batchProcessor) processBatch(ctx context.Context, userCount, itemMinCount, itemMaxCount int) error {
	customerID := bp.generateCustomerID()
	userIDs := bp.generateUserIDs(userCount)
	slog.InfoContext(ctx, "gen redis customer key", "customer_id", customerID)

	for _, userID := range userIDs {
		userItemIDs := bp.generateItemIDs(itemMinCount, itemMaxCount)

		key := bp.formatRedisKey(customerID, userID)
		itemIDStrings := make([]string, len(userItemIDs))

		for i := range len(userItemIDs) {
			itemID := userItemIDs[i]
			itemIDStrings[i] = fmt.Sprintf("%d", itemID)
		}

		bp.userItemMap[key] = strings.Join(itemIDStrings, ",")
	}

	expiration := 24 * time.Hour
	if err := bp.redisClient.SetBatchWithPipeline(ctx, bp.userItemMap, expiration); err != nil {
		slog.ErrorContext(ctx, "failed to save data to Redis", "error", err)
		return fmt.Errorf("failed to save data to Redis: %w", err)
	}

	return nil
}
