package main

import (
	"log"
)

func main() {
	if err := handler(); err != nil {
		log.Fatal(err)
	}
}
