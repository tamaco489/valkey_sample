package main

// Customer ID constants
const (
	CustomerIDMin = 1001
	CustomerIDMax = 1999
)

// User ID constants
const (
	UserIDStart = 10000001
	UserIDEnd   = 10100000
)

// Item ID constants
const (
	ItemIDStart = 30000001
	ItemIDEnd   = 30100000
)

// Data generation mode constants
const (
	// 大量データ用設定
	largeDataUserCount    = 1000 // ランダムに選択する user 数
	largeDataItemMinCount = 100  // 最小選択 item 数
	largeDataItemMaxCount = 500  // 最大選択 item 数

	// 少量データ用設定
	smallDataUserCount    = 3 // ランダムに選択する user 数
	smallDataItemMinCount = 2 // 最小選択 item 数
	smallDataItemMaxCount = 5 // 最大選択 item 数
)

// Redis key format constants
const (
	RedisKeyFormat = "customer:{%d}:user:{%d}:items"
)

// Data format constants
const (
	DataFormat = "user_item_data_%d_%d_%d"
)
