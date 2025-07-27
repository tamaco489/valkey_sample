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

// setUserItemData sets user-item data
func (bp *batchProcessor) setUserItemData(ctx context.Context, customerID, userID, itemID uint32) error {
	key := bp.formatRedisKey(customerID, userID)

	data := fmt.Sprintf(DataFormat, customerID, userID, itemID)

	slog.InfoContext(ctx, "Setting user item data",
		"key", key,
		"data", data,
		"customerID", customerID,
		"userID", userID,
		"itemID", itemID)

	return nil
}

// ProcessBatch processes the entire batch
func (bp *batchProcessor) processBatch(ctx context.Context, userCount, itemMinCount, itemMaxCount int) error {
	userIDs := bp.generateUserIDs(userCount)
	itemIDs := bp.generateItemIDs(itemMinCount, itemMaxCount)

	slog.InfoContext(ctx, "Generated IDs",
		"userCount", len(userIDs),
		"itemCount", len(itemIDs))

	var totalProcessed int
	var totalErrors int

	// 各userIDに対して、ランダムに選択されたitemIDを関連付ける
	for _, userID := range userIDs {
		customerID := bp.generateCustomerID()

		// このuserIDに対して、ランダムに選択されたitemIDを選択
		userItemCount := rand.Intn(itemMaxCount-itemMinCount+1) + itemMinCount
		userItemIDs := bp.generateItemIDs(itemMinCount, itemMaxCount) // 新しいランダム選択

		// 選択されたitemIDの数だけ処理
		for i := range userItemCount {
			if i >= len(userItemIDs) {
				break
			}
			itemID := userItemIDs[i]

			err := bp.setUserItemData(ctx, customerID, userID, itemID)
			if err != nil {
				slog.ErrorContext(ctx, "Failed to set user item data",
					"error", err,
					"customerID", customerID,
					"userID", userID,
					"itemID", itemID)
				totalErrors++
				continue
			}

			totalProcessed++
		}
	}

	slog.InfoContext(ctx, "Batch processing completed",
		"totalProcessed", totalProcessed,
		"totalErrors", totalErrors)

	return nil
}

func handler(isLargeData bool) error {
	ctx := context.Background()

	// データサイズに応じた設定を決定
	var userCount, itemMinCount, itemMaxCount int
	if isLargeData {
		userCount = largeDataUserCount
		itemMinCount = largeDataItemMinCount
		itemMaxCount = largeDataItemMaxCount
	} else {
		userCount = smallDataUserCount
		itemMinCount = smallDataItemMinCount
		itemMaxCount = smallDataItemMaxCount
	}

	slog.InfoContext(ctx, "Starting batch processing",
		"dataSize", map[bool]string{true: "large", false: "small"}[isLargeData],
		"userCount", userCount,
		"itemMinCount", itemMinCount,
		"itemMaxCount", itemMaxCount)

	batchProcessor := newBatchProcessor()

	if err := batchProcessor.processBatch(ctx, userCount, itemMinCount, itemMaxCount); err != nil {
		return fmt.Errorf("batch processing failed: %w", err)
	}

	slog.InfoContext(ctx, "Batch processing completed successfully")
	return nil
}
