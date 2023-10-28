package client

import (
	"context"
	"github.com/redis/go-redis/v9"
	"requestLimiter/config"
	"time"
)

type RedisClient struct {
	conn *redis.Client
}

func NewRedisClient(cfg *config.Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &RedisClient{client}
}

func (r *RedisClient) Add(ctx context.Context, userID string, timeStamp time.Time) error {
	return r.conn.ZAdd(ctx, userID, redis.Z{Score: float64(timeStamp.Unix()), Member: timeStamp}).Err()
}

func (r *RedisClient) Count(ctx context.Context, userID string, from, to string) (int64, error) {
	res := r.conn.ZCount(ctx, userID, from, to)
	if res.Err() != nil {
		return 0, res.Err()
	}
	return res.Val(), nil
}
