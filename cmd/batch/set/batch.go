package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
)

// batchProcessor handles batch operations
type batchProcessor struct {
}

// newBatchProcessor creates a new batch processor
func newBatchProcessor() *batchProcessor {
	return &batchProcessor{}
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

		for i := range userItemCount {
			if i >= len(userItemIDs) {
				break
			}
			itemID := userItemIDs[i]
			key := bp.formatRedisKey(customerID, userID)
			slog.InfoContext(ctx, "setting user item data", "key", key, "item_id", itemID)
		}
	}

	return nil
}
