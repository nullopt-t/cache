package main

import (
	"context"
	"redis-in-practice/cache"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rds := cache.NewRedisCache(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	defer rds.Close()

	ctx := context.Background()

	err := rds.Set(ctx, "key", "value", time.Millisecond*20)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond * 10)

	val, err := rds.Get(ctx, "key")
	if err != nil {
		panic(err)
	}

	println(val)
}
