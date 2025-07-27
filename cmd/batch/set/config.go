package main

// Data generation configuration
type DataConfig struct {
	UserCount    int
	ItemMinCount int
	ItemMaxCount int
}

// GetDataConfig returns configuration based on data size flag
func GetDataConfig(isLargeData bool) DataConfig {
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
