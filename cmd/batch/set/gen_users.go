package main

import "math/rand"

// generateUserIDs generates random user IDs in the specified range
func (bp *batchProcessor) generateUserIDs() []uint32 {
	// 全範囲のuserIDを生成
	allUserIDs := make([]uint32, UserIDEnd-UserIDStart+1)
	for i := range UserIDEnd - UserIDStart + 1 {
		allUserIDs[i] = uint32(UserIDStart + i)
	}

	// ランダムに1000個選択
	selectedUserIDs := make([]uint32, UserIDCount)
	rand.Shuffle(len(allUserIDs), func(i, j int) {
		allUserIDs[i], allUserIDs[j] = allUserIDs[j], allUserIDs[i]
	})

	for i := range UserIDCount {
		selectedUserIDs[i] = allUserIDs[i]
	}

	return selectedUserIDs
}
