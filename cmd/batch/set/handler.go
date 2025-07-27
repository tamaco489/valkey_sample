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

	// データの例（実際の要件に応じて変更）
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
func (bp *batchProcessor) processBatch(ctx context.Context) error {
	slog.InfoContext(ctx, "Starting batch processing")

	userIDs := bp.generateUserIDs()
	itemIDs := bp.generateItemIDs()

	slog.InfoContext(ctx, "Generated IDs",
		"userCount", len(userIDs),
		"itemCount", len(itemIDs))

	var totalProcessed int
	var totalErrors int

	// 各userIDに対して、ランダムに選択されたitemIDを関連付ける
	for _, userID := range userIDs {
		customerID := bp.generateCustomerID()

		// このuserIDに対して、ランダムに100-500個のitemIDを選択
		userItemCount := rand.Intn(ItemIDMaxCount-ItemIDMinCount+1) + ItemIDMinCount
		userItemIDs := bp.generateItemIDs() // 新しいランダム選択

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

func handler() error {
	ctx := context.Background()

	// バッチプロセッサーを作成
	batchProcessor := newBatchProcessor()

	// バッチ処理を実行
	err := batchProcessor.processBatch(ctx)
	if err != nil {
		return fmt.Errorf("batch processing failed: %w", err)
	}

	slog.InfoContext(ctx, "Batch processing completed successfully")
	return nil
}
