package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/tamaco489/valkey_sample/internal/config"
	"github.com/tamaco489/valkey_sample/pkg/core/infrastructure/redis"
)

// batchProcessor handles batch operations
type batchProcessor struct {
	userItemMap map[string]string // キー: customer:{id}:user:{id}:items, 値: item_idのリスト
	redisClient redis.RedisService
}

// newBatchProcessor creates a new batch processor
func newBatchProcessor() *batchProcessor {
	// 設定を読み込み
	cfg := config.Load()

	// Redis接続情報を構築
	redisAddr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	redisDB, _ := strconv.Atoi(cfg.Redis.DB)

	// Redisクライアントを作成
	redisClient, err := redis.NewClient(redisAddr, cfg.Redis.Password, redisDB)
	if err != nil {
		// エラーハンドリングは簡略化（実際の実装では適切に処理）
		slog.Error("failed to create Redis client", "error", err)
		return &batchProcessor{
			userItemMap: make(map[string]string),
		}
	}

	return &batchProcessor{
		userItemMap: make(map[string]string),
		redisClient: redisClient,
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

		bp.userItemMap[key] = strings.Join(itemIDStrings, ",")
	}

	slog.InfoContext(ctx, "user item map contents", "total_keys", len(bp.userItemMap))
	for key, itemIDs := range bp.userItemMap {
		slog.InfoContext(ctx, "user items", "key", key, "item_count", len(itemIDs), "item_ids", itemIDs)
	}

	// Redis に pipelineで データを保存
	if bp.redisClient != nil {
		slog.InfoContext(ctx, "saving data to Redis using pipeline")
		if err := bp.redisClient.SetWithPipeline(ctx, bp.userItemMap, 24*time.Hour); err != nil {
			slog.ErrorContext(ctx, "failed to save data to Redis", "error", err)
			return fmt.Errorf("failed to save data to Redis: %w", err)
		}
		slog.InfoContext(ctx, "successfully saved data to Redis using pipeline")
	}

	bp.redisClient.Close()

	return nil
}
