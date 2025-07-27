package main

import (
	"context"
	"fmt"
	"log/slog"
)

// Data generation configuration
type DataConfig struct {
	UserCount    int
	ItemMinCount int
	ItemMaxCount int
}

// NewDataConfig creates a new data configuration
// Default is small data, use isLarge=true for large data
func NewDataConfig(isLarge bool) DataConfig {
	if isLarge {
		return DataConfig{
			UserCount:    largeDataUserCount,
			ItemMinCount: largeDataItemMinCount,
			ItemMaxCount: largeDataItemMaxCount,
		}
	}
	return DataConfig{
		UserCount:    smallDataUserCount,
		ItemMinCount: smallDataItemMinCount,
		ItemMaxCount: smallDataItemMaxCount,
	}
}

func handler(isLarge bool) error {
	ctx := context.Background()
	cfg := NewDataConfig(isLarge)
	slog.InfoContext(ctx, "Starting batch processing",
		"dataSize", map[bool]string{true: "large", false: "small"}[isLarge],
		"userCount", cfg.UserCount,
		"itemMinCount", cfg.ItemMinCount,
		"itemMaxCount", cfg.ItemMaxCount)

	batchProcessor := newBatchProcessor()

	if err := batchProcessor.processBatch(ctx, cfg.UserCount, cfg.ItemMinCount, cfg.ItemMaxCount); err != nil {
		return fmt.Errorf("batch processing failed: %w", err)
	}

	slog.InfoContext(ctx, "Batch processing completed successfully")
	return nil
}
