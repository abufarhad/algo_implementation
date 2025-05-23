package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password for local development
		DB:       0,                // Default DB
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	// Set a key-value pair
	sErr := client.Set(ctx, "greeting", "Hello, Redis!", 0).Err()
	if sErr != nil {
		log.Fatal(err)
	}
	// Get the value for a key
	value, err := client.Get(ctx, "greeting").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Greeting:", value)

	// working with list

	// LPush nsert values at the head of the list
	err = client.LPush(ctx, "tasks", "Task 1", "Task 2").Err()
	if err != nil {
		log.Fatal(err)
	}

	// RPush insert values at the back/ tail of the list
	err = client.RPush(ctx, "tasks", "Task 1", "Task 2").Err()
	if err != nil {
		log.Fatal(err)
	}
	// LPop remove the first element in the list
	task, err := client.LPop(ctx, "tasks").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Popped Task:", task)

	// RPopr remove the last element in the list
	taskRpop, err := client.RPop(ctx, "tasks").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Popped Task:", taskRpop)

}
