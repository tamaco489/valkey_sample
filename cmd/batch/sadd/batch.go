package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/tamaco489/valkey_sample/internal/config"
	"github.com/tamaco489/valkey_sample/pkg/core/infrastructure/redis"
)

// batchProcessor handles batch operations
type batchProcessor struct {
	userItemMap map[string][]string // キー: customer:{customer_id}:user:{user_id}:items, 値: item_idのスライス
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
		os.Exit(1)
	}

	return &batchProcessor{
		userItemMap: make(map[string][]string),
		redisClient: redisClient,
	}
}

// formatRedisKey creates a Redis key with the specified format
func (bp *batchProcessor) formatRedisKey(customerID, userID uint32) string {
	return fmt.Sprintf(RedisKeyFormat, customerID, userID)
}

// generateCustomerID generates a random customer ID
func (bp *batchProcessor) generateCustomerID() uint32 {
	return uint32(rand.Intn(CustomerIDMax-CustomerIDMin+1) + CustomerIDMin)
}

// generateUserIDs generates random user IDs in the specified range
func (bp *batchProcessor) generateUserIDs(userCount int) []uint32 {
	allUserIDs := make([]uint32, UserIDEnd-UserIDStart+1)
	for i := range UserIDEnd - UserIDStart + 1 {
		allUserIDs[i] = uint32(UserIDStart + i)
	}

	selectedUserIDs := make([]uint32, userCount)
	rand.Shuffle(len(allUserIDs), func(i, j int) {
		allUserIDs[i], allUserIDs[j] = allUserIDs[j], allUserIDs[i]
	})

	for i := range userCount {
		selectedUserIDs[i] = allUserIDs[i]
	}

	return selectedUserIDs
}

// generateItemIDs generates random item IDs in the specified range
func (bp *batchProcessor) generateItemIDs(itemMinCount, itemMaxCount int) []uint32 {
	allItemIDs := make([]uint32, ItemIDEnd-ItemIDStart+1)
	for i := range ItemIDEnd - ItemIDStart + 1 {
		allItemIDs[i] = uint32(ItemIDStart + i)
	}

	rand.Shuffle(len(allItemIDs), func(i, j int) {
		allItemIDs[i], allItemIDs[j] = allItemIDs[j], allItemIDs[i]
	})

	itemCount := rand.Intn(itemMaxCount-itemMinCount+1) + itemMinCount
	selectedItemIDs := make([]uint32, itemCount)

	for i := range itemCount {
		selectedItemIDs[i] = allItemIDs[i]
	}

	return selectedItemIDs
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

		bp.userItemMap[key] = itemIDStrings
	}

	slog.InfoContext(ctx, "user item map contents", "total_keys", len(bp.userItemMap))
	for key, itemIDs := range bp.userItemMap {
		slog.InfoContext(ctx, "user items", "key", key, "item_count", len(itemIDs), "item_ids", itemIDs)
	}

	expiration := 24 * time.Hour
	if err := bp.redisClient.SAddBatchWithPipeline(ctx, bp.userItemMap, expiration); err != nil {
		slog.ErrorContext(ctx, "failed to save data to Redis", "error", err)
		return fmt.Errorf("failed to save data to Redis: %w", err)
	}

	return nil
}
