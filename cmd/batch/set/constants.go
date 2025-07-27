package main

// Customer ID constants
const (
	CustomerIDMin = 1001
	CustomerIDMax = 1999
)

// User ID constants
const (
	UserIDStart = 10000001
	UserIDEnd   = 10001000
)

// Item ID constants
const (
	ItemIDStart = 30000001
	ItemIDEnd   = 30001000
)

// Redis key format constants
const (
	RedisKeyFormat = "customer:%d:user:%d:items"
)

// Data format constants
const (
	DataFormat = "user_item_data_%d_%d_%d"
)
