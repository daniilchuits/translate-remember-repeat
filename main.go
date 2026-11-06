package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("Error loading '.env': %s", err))
	}

	cnnStr := os.Getenv("REDIS_CNN_STR")
	fmt.Println(cnnStr)
	client := redis.NewClient(&redis.Options{
		Addr: cnnStr,
	})

	client.Set(context.Background(), "hello", "world", time.Second)
	fmt.Println(client.Get(context.Background(), "hello"))
}
