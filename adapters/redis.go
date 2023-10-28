package adapters

import (
	"context"
	"time"
)

type RedisAdapter interface {
	Add(ctx context.Context, userID string, timeStamp time.Time) error
	Count(ctx context.Context, userID string, from, to string) (int64, error)
}
