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
	LargeDataUserCount    = 1000 // ランダムに選択するuser数
	LargeDataItemMinCount = 100  // 最小選択item数
	LargeDataItemMaxCount = 500  // 最大選択item数

	// 少量データ用設定
	SmallDataUserCount    = 3 // ランダムに選択するuser数
	SmallDataItemMinCount = 2 // 最小選択item数
	SmallDataItemMaxCount = 5 // 最大選択item数
)

// Redis key format constants
const (
	RedisKeyFormat = "customer:%d:user:%d:items"
)

// Data format constants
const (
	DataFormat = "user_item_data_%d_%d_%d"
)
