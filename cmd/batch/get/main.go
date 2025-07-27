package main

import (
	"flag"
	"log"
)

func main() {
	var customerID uint
	flag.UintVar(&customerID, "customer_id", 0, "Customer ID to fetch data for (required)")
	flag.Parse()

	if customerID == 0 {
		log.Fatal("customer_id is required. Usage: -customer_id=1234")
	}

	if err := handler(uint32(customerID)); err != nil {
		log.Fatal(err)
	}
}
