package main

import "math/rand"

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
