package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(
	opt *redis.Options,
) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(opt),
	}
}

func (c *RedisCache) Get(
	ctx context.Context,
	key string,
) (string, error) {
	cmd := c.client.Get(ctx, key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}

func (c *RedisCache) Set(
	ctx context.Context,
	key,
	value string,
	expiration time.Duration,
) error {
	cmd := c.client.Set(ctx, key, value, expiration)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (c *RedisCache) Delete(
	ctx context.Context,
	key string,
) error {
	cmd := c.client.Del(ctx, key)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (c *RedisCache) Close() error {
	return c.client.Close()
}
