package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var ctx = context.Background()

func main() {
	// Initialize Redis Cluster client
	// Include the addresses of all Redis instances (both masters and replicas)
	// Allows the client to discover the cluster topology, handle master-slave failover, and distribute commands appropriately across the cluster
	// Get Redis addresses from environment variable
	addrs := strings.Split(os.Getenv("REDIS_ADDRESSES"), ",")

	// Initialize Redis Cluster client
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        addrs,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// To iterate over shards:
	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		log.Fatalf("Failed to ping shards: %v", err)
	}

	// Sleep for a few seconds to allow the cluster to stabilize
	time.Sleep(5 * time.Second)
	fmt.Printf("Redis Cluster initialized successfully\n")

	contestIDs := []string{"contest:12345", "contest:67890", "contest:54321"}

	for _, contestID := range contestIDs {
		// Add scores for 10 users
		scores := []float64{10, 30, 100, 50, 70, 90, 20, 80, 40, 60}

		for i, score := range scores {
			userID := fmt.Sprintf("user:%d", i)
			err := rdb.ZAdd(ctx, contestID, redis.Z{Score: score, Member: userID}).Err()
			if err != nil {
				log.Fatalf("Failed to add score: %v", err)
			}
		}

		// Fetch top 2 users
		topUsers, err := rdb.ZRevRangeWithScores(ctx, contestID, 0, 1).Result()
		if err != nil {
			log.Fatalf("Failed to fetch leaderboard: %v", err)
		}

		fmt.Println("Top 2 Users:")
		for i, z := range topUsers {
			fmt.Printf("%d. %s - Score: %f\n", i+1, z.Member, z.Score)
		}
	}

	// Prevent the application from exiting
	log.Println("Go application running. Press CTRL+C to exit.")
	select {} // Block forever
}
