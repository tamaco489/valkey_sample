package main

import (
	"context"
	"log/slog"
)

func handler(customerID uint32) error {
	ctx := context.Background()
	slog.InfoContext(ctx, "Starting GET batch processing", "customerID", customerID)

	batchProcessor := newBatchProcessor()

	if err := batchProcessor.processBatch(ctx, customerID); err != nil {
		return err
	}

	slog.InfoContext(ctx, "GET batch processing completed successfully")
	return nil
}
