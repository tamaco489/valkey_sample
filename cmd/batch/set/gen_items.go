package main

import "math/rand"

// generateItemIDs generates random item IDs in the specified range
func (bp *batchProcessor) generateItemIDs() []uint32 {
	// 全範囲のitemIDを生成
	allItemIDs := make([]uint32, ItemIDEnd-ItemIDStart+1)
	for i := range ItemIDEnd - ItemIDStart + 1 {
		allItemIDs[i] = uint32(ItemIDStart + i)
	}

	// ランダムに100-500個選択
	itemCount := rand.Intn(ItemIDMaxCount-ItemIDMinCount+1) + ItemIDMinCount
	selectedItemIDs := make([]uint32, itemCount)
	rand.Shuffle(len(allItemIDs), func(i, j int) {
		allItemIDs[i], allItemIDs[j] = allItemIDs[j], allItemIDs[i]
	})

	for i := range itemCount {
		selectedItemIDs[i] = allItemIDs[i]
	}

	return selectedItemIDs
}
