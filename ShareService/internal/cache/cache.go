package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

type Cache interface {
	Get(key string, dataChan chan<- string, errCh chan<- error)
	Cache(key string, value string, dur time.Duration, errCh chan<- error)
	Delete(key string, errCh chan<- error)
}

func NewRedisCache(address string, password string) (*redisCache, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})

	err := c.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &redisCache{
		client: c,
	}, nil
}

func (c *redisCache) Get(key string, dataCh chan<- string, errCh chan<- error) {
	res, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
	dataCh <- res
}

func (c *redisCache) Cache(key string, value string, dur time.Duration, errCh chan<- error) {
	err := c.client.Set(context.Background(), key, value, dur).Err()
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}

func (c *redisCache) Delete(key string, errCh chan<- error) {
	err := c.client.Del(context.Background(), key).Err()
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}
