package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
)

// batchProcessor handles batch operations
type batchProcessor struct {
	userItemMap map[string][]string // キー: customer:{id}:user:{id}:items, 値: item_idのリスト
}

// newBatchProcessor creates a new batch processor
func newBatchProcessor() *batchProcessor {
	return &batchProcessor{
		userItemMap: make(map[string][]string),
	}
}

// formatRedisKey creates a Redis key with the specified format
func (bp *batchProcessor) formatRedisKey(customerID, userID uint32) string {
	return fmt.Sprintf(RedisKeyFormat, customerID, userID)
}

// processBatch processes the entire batch
func (bp *batchProcessor) processBatch(ctx context.Context, userCount, itemMinCount, itemMaxCount int) error {
	userIDs := bp.generateUserIDs(userCount)
	itemIDs := bp.generateItemIDs(itemMinCount, itemMaxCount)
	slog.InfoContext(ctx, "generated ids", "user_count", len(userIDs), "item_count", len(itemIDs))

	for _, userID := range userIDs {
		customerID := bp.generateCustomerID()
		userItemCount := rand.Intn(itemMaxCount-itemMinCount+1) + itemMinCount
		userItemIDs := bp.generateItemIDs(itemMinCount, itemMaxCount)

		key := bp.formatRedisKey(customerID, userID)
		itemIDStrings := make([]string, userItemCount)

		for i := range userItemCount {
			if i >= len(userItemIDs) {
				break
			}
			itemID := userItemIDs[i]
			itemIDStrings[i] = fmt.Sprintf("%d", itemID)
		}
		bp.userItemMap[key] = itemIDStrings
	}

	slog.InfoContext(ctx, "user item map contents", "total_keys", len(bp.userItemMap))
	for key, itemIDs := range bp.userItemMap {
		slog.InfoContext(ctx, "user items", "key", key, "item_count", len(itemIDs), "item_ids", itemIDs)
	}

	return nil
}
