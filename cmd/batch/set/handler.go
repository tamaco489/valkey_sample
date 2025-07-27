package main

import (
	"context"
	"fmt"
	"log/slog"
)

func handler(isLargeData bool) error {
	ctx := context.Background()
	cfg := NewDataConfig(isLargeData)
	slog.InfoContext(ctx, "Starting batch processing",
		"dataSize", map[bool]string{true: "large", false: "small"}[isLargeData],
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
