package storage

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client}
}

func (store *RedisStore) Get(ctx context.Context, code string) (string, error) {
	value, err := store.client.Get(ctx, code).Result()
	if err == nil {
		return value, nil
	}
	if errors.Is(err, redis.Nil) {
		return "", NotFoundError{Code: code}
	}
	return "", err
}

func (store *RedisStore) Set(ctx context.Context, code string, target string) error {
	return store.client.Set(ctx, code, target, 0).Err()
}

func (store *RedisStore) SetIfAbsent(ctx context.Context, code string, target string) (bool, error) {
	ok, err := store.client.SetNX(ctx, code, target, 0).Result()
	return ok, err
}
