package cache

import (
	"errors"

	"github.com/go-redis/redis"
)

type redisCache struct {
	cache *redis.Client
}

func NewRedisCache(cache *redis.Client) Cacher {
	return &redisCache{
		cache: cache,
	}
}

func (c *redisCache) Get(key string) ([]byte, error) {
	obj := c.cache.Get(key)
	val, err := obj.Result()
	if err != nil {
		return []byte{}, errors.New("We could not get cache")
	}
	return []byte(val), nil
}
func (c *redisCache) Set(key string, object []byte) error {
	return nil
}
