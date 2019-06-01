package main

import (
	"github.com/go-redis/redis"
	"github.com/nozo-moto/google_like_search_engine/crawler"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password",
	})

	c := crawler.New(*redisClient)
	c.Run()
}
