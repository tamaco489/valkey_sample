package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
)

// BatchProcessor handles batch operations
type BatchProcessor struct {
}

// NewBatchProcessor creates a new batch processor
func NewBatchProcessor() *BatchProcessor {
	return &BatchProcessor{}
}

// GenerateUserIDs generates user IDs in the specified range
func (bp *BatchProcessor) GenerateUserIDs() []uint32 {
	var userIDs []uint32
	for i := UserIDStart; i <= UserIDEnd; i++ {
		userIDs = append(userIDs, uint32(i))
	}
	return userIDs
}

// GenerateItemIDs generates item IDs in the specified range
func (bp *BatchProcessor) GenerateItemIDs() []uint32 {
	var itemIDs []uint32
	for i := ItemIDStart; i <= ItemIDEnd; i++ {
		itemIDs = append(itemIDs, uint32(i))
	}
	return itemIDs
}

// GenerateCustomerID generates a random customer ID
func (bp *BatchProcessor) GenerateCustomerID() uint32 {
	return uint32(rand.Intn(CustomerIDMax-CustomerIDMin+1) + CustomerIDMin)
}

// CreateRedisKey creates a Redis key with the specified format
func (bp *BatchProcessor) CreateRedisKey(customerID, userID uint32) string {
	return fmt.Sprintf(RedisKeyFormat, customerID, userID)
}

// SetUserItemData sets user-item data
func (bp *BatchProcessor) SetUserItemData(ctx context.Context, customerID, userID, itemID uint32) error {
	key := bp.CreateRedisKey(customerID, userID)

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
func (bp *BatchProcessor) ProcessBatch(ctx context.Context) error {
	slog.InfoContext(ctx, "Starting batch processing")

	userIDs := bp.GenerateUserIDs()
	itemIDs := bp.GenerateItemIDs()

	totalProcessed := 0
	totalErrors := 0

	for _, userID := range userIDs {
		for _, itemID := range itemIDs {
			customerID := bp.GenerateCustomerID()

			err := bp.SetUserItemData(ctx, customerID, userID, itemID)
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
	batchProcessor := NewBatchProcessor()

	// バッチ処理を実行
	err := batchProcessor.ProcessBatch(ctx)
	if err != nil {
		return fmt.Errorf("batch processing failed: %w", err)
	}

	slog.InfoContext(ctx, "Batch processing completed successfully")
	return nil
}
