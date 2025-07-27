package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

// Example usage of Redis client
func Example() {
	// Load configuration
	config := LoadConfig()
	
	// Create factory
	factory := NewFactory(config)
	
	// Create client
	client, err := factory.CreateClient()
	if err != nil {
		slog.Error("Failed to create Redis client", "error", err)
		return
	}
	defer client.Close()

	ctx := context.Background()

	// Set a value
	err = client.Set(ctx, "example_key", "example_value", 10*time.Minute)
	if err != nil {
		slog.Error("Failed to set value", "error", err)
		return
	}

	// Get a value
	value, err := client.Get(ctx, "example_key")
	if err != nil {
		slog.Error("Failed to get value", "error", err)
		return
	}

	fmt.Printf("Retrieved value: %s\n", value)

	// Check if key exists
	exists, err := client.Exists(ctx, "example_key")
	if err != nil {
		slog.Error("Failed to check key existence", "error", err)
		return
	}

	fmt.Printf("Key exists: %d\n", exists)

	// Delete key
	err = client.Del(ctx, "example_key")
	if err != nil {
		slog.Error("Failed to delete key", "error", err)
		return
	}

	fmt.Println("Key deleted successfully")
}

// ExampleWithCustomConfig demonstrates using Redis client with custom configuration
func ExampleWithCustomConfig() {
	// Create factory with custom config
	config := &Config{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
		PoolSize: 10,
	}
	
	factory := NewFactory(config)
	
	// Create client with custom configuration
	client, err := factory.CreateClientWithConfig("localhost:6379", "", 0)
	if err != nil {
		slog.Error("Failed to create Redis client with custom config", "error", err)
		return
	}
	defer client.Close()

	ctx := context.Background()

	// Test ping
	err = client.Ping(ctx)
	if err != nil {
		slog.Error("Failed to ping Redis", "error", err)
		return
	}

	fmt.Println("Redis connection successful")
} 