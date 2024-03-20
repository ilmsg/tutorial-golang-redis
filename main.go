package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// println("Hello")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	r := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0, // default DB
	})

	ctx := context.Background()
	valCount, err := r.Get(ctx, "count").Result()
	if err == redis.Nil {
		valCount = "0"
	} else if err != nil {
		log.Fatal(err)
	}

	count, _ := strconv.Atoi(valCount)
	count++
	if err := r.Set(ctx, "count", count, 0).Err(); err != nil {
		log.Fatal(err)
	}
	println("count:", count)
}
