package cache

import "github.com/go-redis/redis"

type Cacher interface {
	Get(key string) ([]byte, error)
	Set(key string, object []byte) error
}

func RedisHandler(addr string, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
