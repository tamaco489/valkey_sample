package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/tamaco489/valkey_sample/pkg/core/infrastructure/redis"
)

// batchProcessor handles batch operations
type batchProcessor struct {
	userItemMap map[string][]uint32 // キー: customer:{customer_id}:user:{user_id}:items, 値: item_idのスライス
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
		userItemMap: make(map[string][]uint32),
		redisClient: redisClient,
	}
}

// formatRedisKey creates a Redis key with the specified format
func (bp *batchProcessor) formatRedisKey(customerID, userID uint32) string {
	return fmt.Sprintf(RedisKeyFormat, customerID, userID)
}

// generateAllUserIDs generates all user IDs in the range
func (bp *batchProcessor) generateAllUserIDs() []uint32 {
	totalUsers := UserIDEnd - UserIDStart + 1
	userIDs := make([]uint32, totalUsers)

	for i := range totalUsers {
		userIDs[i] = uint32(UserIDStart + i)
	}

	return userIDs
}

// parseItemIDsFromValue parses comma-separated item IDs from Redis value
func (bp *batchProcessor) parseItemIDsFromValue(value string) ([]uint32, error) {
	if value == "" {
		return []uint32{}, nil
	}

	itemIDStrings := strings.Split(value, ",")
	itemIDs := make([]uint32, 0, len(itemIDStrings))

	for _, idStr := range itemIDStrings {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}

		itemID, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse item ID %s: %w", idStr, err)
		}
		itemIDs = append(itemIDs, uint32(itemID))
	}

	return itemIDs, nil
}

// processBatch processes the entire batch
func (bp *batchProcessor) processBatch(ctx context.Context, customerID uint32) error {
	userIDs := bp.generateAllUserIDs()
	slog.InfoContext(ctx, "using specified customer ID", "customer_id", customerID)

	keys := make([]string, len(userIDs))
	for i, userID := range userIDs {
		keys[i] = bp.formatRedisKey(customerID, userID)
	}

	redisData, err := bp.redisClient.GetBatchWithPipeline(ctx, keys)
	if err != nil {
		return fmt.Errorf("failed to fetch data from Redis: %w", err)
	}

	for key, value := range redisData {
		itemIDs, err := bp.parseItemIDsFromValue(value)
		if err != nil {
			return fmt.Errorf("failed to parse item IDs for key %s: %w", key, err)
		}
		bp.userItemMap[key] = itemIDs
	}

	var count int
	for key, itemIDs := range bp.userItemMap {
		if count >= 5 {
			slog.InfoContext(ctx, "... and more data (showing first 5 entries only)")
			break
		}
		slog.InfoContext(ctx, "user items", "key", key, "item_count", len(itemIDs), "item_ids", itemIDs)
		count++
	}

	return nil
}
