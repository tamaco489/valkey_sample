package main

import (
	"flag"
	"log"
)

func main() {
	var isLargeData bool
	flag.BoolVar(&isLargeData, "large", false, "Generate large amount of data (default: small)")
	flag.Parse()

	if err := handler(isLargeData); err != nil {
		log.Fatal(err)
	}
}
