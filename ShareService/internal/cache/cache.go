package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache(address string, password string) (*Cache, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})

	err := c.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &Cache{
		client: c,
	}, nil
}

func (c *Cache) Get(key string, cacheCh chan<- string, errCh chan<- error) {
	res, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
	cacheCh <- res
}

func (c *Cache) Cache(key string, value string, dur time.Duration, errCh chan<- error) {
	err := c.client.Set(context.Background(), key, value, dur).Err()
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}

func (c *Cache) Delete(key string, errCh chan<- error) {
	err := c.client.Del(context.Background(), key).Err()
	if err != nil {
		errCh <- err
		return
	}

	errCh <- nil
}
