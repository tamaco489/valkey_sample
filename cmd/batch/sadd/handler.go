package main

import (
	"context"
	"log/slog"
)

func handler(isLarge bool) error {
	ctx := context.Background()
	cfg := NewDataConfig(isLarge)
	slog.InfoContext(ctx, "Starting SADD batch processing",
		"dataSize", map[bool]string{true: "large", false: "small"}[isLarge],
		"userCount", cfg.UserCount,
		"itemMinCount", cfg.ItemMinCount,
		"itemMaxCount", cfg.ItemMaxCount)

	batchProcessor := newBatchProcessor()

	if err := batchProcessor.processBatch(ctx, cfg.UserCount, cfg.ItemMinCount, cfg.ItemMaxCount); err != nil {
		return err
	}

	slog.InfoContext(ctx, "SADD batch processing completed successfully")
	return nil
}
