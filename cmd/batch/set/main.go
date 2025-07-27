package main

import (
	"context"
	"log"
	"log/slog"
)

func handler() error {
	ctx := context.Background()
	slog.InfoContext(ctx, "start redis set")
	return nil
}

func main() {
	if err := handler(); err != nil {
		log.Fatal(err)
	}
}
