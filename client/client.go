package client

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"requestLimiter/config"
	"time"
)

type RedisClient struct {
	conn *redis.Client
}

func NewRedisClient(cfg *config.Config) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	status := client.Ping(context.Background())
	if status.Err() != nil {
		return nil, status.Err()
	}
	return &RedisClient{client}, nil
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
