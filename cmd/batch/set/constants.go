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
	UserIDCount = 1000 // ランダムに選択する数
	// UserIDCount = 3 // ランダムに選択する数 (少量のデータで確認する時用)
)

// Item ID constants
const (
	ItemIDStart    = 30000001
	ItemIDEnd      = 30100000
	ItemIDMinCount = 100 // 最小選択数
	ItemIDMaxCount = 500 // 最大選択数
)

// Redis key format constants
const (
	RedisKeyFormat = "customer:%d:user:%d:items"
)

// Data format constants
const (
	DataFormat = "user_item_data_%d_%d_%d"
)
