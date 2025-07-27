package main

// Data generation configuration
type DataConfig struct {
	UserCount    int
	ItemMinCount int
	ItemMaxCount int
}

// NewDataConfig creates a new data configuration
// Default is small data, use isLargeData=true for large data
func NewDataConfig(isLargeData bool) DataConfig {
	if isLargeData {
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
