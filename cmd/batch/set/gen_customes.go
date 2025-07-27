package main

import "math/rand"

// generateCustomerID generates a random customer ID
func (bp *batchProcessor) generateCustomerID() uint32 {
	return uint32(rand.Intn(CustomerIDMax-CustomerIDMin+1) + CustomerIDMin)
}
