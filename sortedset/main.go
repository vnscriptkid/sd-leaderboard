package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()

func main() {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Redis database number
	})

	// Simulate contest leaderboard
	contestID := "contest:12345"

	// Add scores for 10 users
	scores := []float64{10, 30, 100, 50, 70, 90, 20, 80, 40, 60}

	for i, score := range scores {
		userID := fmt.Sprintf("user:%d", i)
		err := rdb.ZAdd(ctx, contestID, &redis.Z{Score: score, Member: userID}).Err()
		if err != nil {
			log.Fatalf("Failed to add score: %v", err)
		}
	}

	// Fetch top 5 users
	topUsers, err := rdb.ZRevRangeWithScores(ctx, contestID, 0, 4).Result()
	if err != nil {
		log.Fatalf("Failed to fetch leaderboard: %v", err)
	}

	fmt.Println("Top 5 Users:")
	for i, z := range topUsers {
		fmt.Printf("%d. %s - Score: %f\n", i+1, z.Member, z.Score)
	}

	// Count users with a score between 50 and 100
	count, err := rdb.ZCount(ctx, contestID, "50", "100").Result()
	if err != nil {
		log.Fatalf("Failed to count users: %v", err)
	}
	fmt.Printf("Number of users with scores between 50 and 100: %d\n", count)
}
