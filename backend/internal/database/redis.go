package database

import (
	"context"
	"os"
	"log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URI"),
		Password: "",
		DB: 0,
	})

	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis. " + err.Error())
	}

	log.Println("Redis connected")
}