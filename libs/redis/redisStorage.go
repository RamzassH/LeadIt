package redisStorage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"time"
)

type Redis struct {
	client *redis.Client
	logger zerolog.Logger
}

func New(client *redis.Client, logger zerolog.Logger) (*Redis, error) {
	if client == nil {
		return nil, fmt.Errorf("redis connection is nil")
	}
	return &Redis{client: client, logger: logger}, nil
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	const op = "redis Set"
	logger := r.logger.With().Str("op", op).Logger()
	logger.Debug().Str("key", key).Interface("value", value).Str("expiration", expiration.String())
	status := r.client.Set(ctx, key, value, expiration)
	logger.Debug().Str("status", status.String()).Msg("set status")
	if status.Err() != nil {
		return fmt.Errorf("%s: %w", op, status.Err())
	}
	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	const op = "redis Get"
	res := r.client.Get(ctx, key)
	if res.Err() != nil {
		return "", fmt.Errorf("%s: %w", op, res.Err())
	}
	return res.Result()
}

func (r *Redis) Del(ctx context.Context, key string) error {
	const op = "redis Del"
	status := r.client.Del(ctx, key)
	if status.Err() != nil {
		return fmt.Errorf("%s: %w", op, status.Err())
	}
	return nil
}

func (r *Redis) HSet(ctx context.Context, key, field string, value interface{}) error {
	const op = "redis HSet"
	status := r.client.HSet(ctx, key, field, value)
	if status.Err() != nil {
		return fmt.Errorf("%s: %w", op, status.Err())
	}
	return nil
}

func (r *Redis) HGet(ctx context.Context, key, field string) (string, error) {
	const op = "redis HGet"
	res := r.client.HGet(ctx, key, field)
	if res.Err() != nil {
		return "", fmt.Errorf("%s: %w", op, res.Err())
	}
	return res.Result()
}

func (r *Redis) HDel(ctx context.Context, key string, fields ...string) error {
	const op = "redis HDel"
	status := r.client.HDel(ctx, key, fields...)
	if status.Err() != nil {
		return fmt.Errorf("%s: %w", op, status.Err())
	}
	return nil
}

func (r *Redis) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	const op = "redis HGetAll"
	res := r.client.HGetAll(ctx, key)
	if res.Err() != nil {
		return nil, fmt.Errorf("%s: %w", op, res.Err())
	}
	return res.Result()
}
