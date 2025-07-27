package main

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
