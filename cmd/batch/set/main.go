package main

import (
	"flag"
	"log"
)

func main() {
	var isLarge bool
	flag.BoolVar(&isLarge, "large", false, "Generate large amount of data (default: small)")
	flag.Parse()

	if err := handler(isLarge); err != nil {
		log.Fatal(err)
	}
}
