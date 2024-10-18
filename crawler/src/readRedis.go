package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

//best link for this: https://upstash.com/docs/redis/tutorials/goapi

func test() {
	client := redis.NewClient(&redis.Options{
		Addr:     "crawlerurlqueue-fvdqpx.serverless.euw1.cache.amazonaws.com:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()
	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	test()
}
