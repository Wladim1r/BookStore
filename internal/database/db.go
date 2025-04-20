package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitDB() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: os.Getenv("DB_PASSWORD"),
		DB:       9,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatal("Could not connect to Redis", err)
	}

	log.Println("Successfuly connected to Redis")

	return redisClient
}
