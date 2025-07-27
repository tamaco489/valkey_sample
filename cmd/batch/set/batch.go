package main

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/tamaco489/valkey_sample/internal/config"
	"github.com/tamaco489/valkey_sample/pkg/core/infrastructure/redis"
)

// batchProcessor handles batch operations
type batchProcessor struct {
	userItemMap map[string]string // キー: customer:{id}:user:{id}:items, 値: item_idのカンマ区切り
	redisClient redis.RedisService
}

// newBatchProcessor creates a new batch processor
func newBatchProcessor() *batchProcessor {
	cfg := config.Load()
	redisAddr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	redisDB, _ := strconv.Atoi(cfg.Redis.DB)
	redisClient, err := redis.NewClient(redisAddr, cfg.Redis.Password, redisDB)
	if err != nil {
		slog.Error("failed to create Redis client", "error", err)
		slog.Info("continuing without Redis connection - data will be generated but not saved")
		return &batchProcessor{
			userItemMap: make(map[string]string),
			redisClient: nil,
		}
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
	if err := bp.redisClient.SetWithPipeline(ctx, bp.userItemMap, expiration); err != nil {
		slog.ErrorContext(ctx, "failed to save data to Redis", "error", err)
		return fmt.Errorf("failed to save data to Redis: %w", err)
	}

	return nil
}
