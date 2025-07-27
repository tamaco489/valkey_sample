package main

import "math/rand"

// generateItemIDs generates random item IDs in the specified range
func (bp *batchProcessor) generateItemIDs(itemMinCount, itemMaxCount int) []uint32 {
	allItemIDs := make([]uint32, ItemIDEnd-ItemIDStart+1)
	for i := range ItemIDEnd - ItemIDStart + 1 {
		allItemIDs[i] = uint32(ItemIDStart + i)
	}

	rand.Shuffle(len(allItemIDs), func(i, j int) {
		allItemIDs[i], allItemIDs[j] = allItemIDs[j], allItemIDs[i]
	})

	itemCount := rand.Intn(itemMaxCount-itemMinCount+1) + itemMinCount
	selectedItemIDs := make([]uint32, itemCount)

	for i := range itemCount {
		selectedItemIDs[i] = allItemIDs[i]
	}

	return selectedItemIDs
}
