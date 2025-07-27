package main

import "math/rand"

// generateCustomerID generates a random customer ID
func (bp *batchProcessor) generateCustomerID() uint32 {
	return uint32(rand.Intn(CustomerIDMax-CustomerIDMin+1) + CustomerIDMin)
}

// generateUserIDs generates random user IDs in the specified range
func (bp *batchProcessor) generateUserIDs(userCount int) []uint32 {
	allUserIDs := make([]uint32, UserIDEnd-UserIDStart+1)
	for i := range UserIDEnd - UserIDStart + 1 {
		allUserIDs[i] = uint32(UserIDStart + i)
	}

	selectedUserIDs := make([]uint32, userCount)
	rand.Shuffle(len(allUserIDs), func(i, j int) {
		allUserIDs[i], allUserIDs[j] = allUserIDs[j], allUserIDs[i]
	})

	for i := range userCount {
		selectedUserIDs[i] = allUserIDs[i]
	}

	return selectedUserIDs
}

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
