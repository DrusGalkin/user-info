package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type App struct {
	client  *redis.Client
	timeout time.Duration
	ttl     time.Duration
}

func New(address string, port int, timeout time.Duration, ttl time.Duration) *App {
	return &App{
		client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", address, port),
		}),
		timeout: timeout,
		ttl:     ttl,
	}
}

func (a *App) GetData(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	result, err := a.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (a *App) SetData(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	set := a.client.Set(ctx, key, value, a.ttl)
	return set.Err()
}
