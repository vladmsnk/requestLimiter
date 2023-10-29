package limiter

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"requestLimiter/adapters"
	"sync"
	"time"
)

type Repo struct {
	rate          int64
	windowSizeSec int64
	mu            sync.Mutex
	rd            adapters.RedisAdapter
}

func NewLimiter(rate, windowSizeSec int64, r adapters.RedisAdapter) *Repo {
	return &Repo{
		rate:          rate,
		windowSizeSec: windowSizeSec,
		rd:            r,
	}
}

type Limiter interface {
	Allow(ctx context.Context, userID string, requestTimeStamp time.Time) (bool, error)
	Limit(ctx context.Context) (bool, error)
}

func (l *Repo) Allow(ctx context.Context, userID string, requestTimeStamp time.Time) (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	from := fmt.Sprintf("%f", float64(requestTimeStamp.Unix())-float64(l.windowSizeSec))
	to := fmt.Sprintf("%f", float64(requestTimeStamp.Unix()))

	totalRequests, err := l.rd.Count(ctx, userID, from, to)
	if err != nil {
		return false, err
	}
	if totalRequests >= l.rate {
		return false, nil
	}

	err = l.rd.Add(ctx, userID, requestTimeStamp)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (l *Repo) Limit(ctx context.Context) (bool, error) {
	requestTime := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, status.Error(codes.Internal, "Failed to extract metadata")
	}

	// Retrieve the user ID from metadata
	userIDs := md.Get("user_id")
	if len(userIDs) == 0 {
		return false, status.Error(codes.Unauthenticated, "User ID not provided")
	}
	userID := userIDs[0]

	allowed, err := l.Allow(ctx, userID, requestTime)
	if err != nil {
		return false, err
	}
	if !allowed {
		return false, status.Error(codes.ResourceExhausted, "Too many requests")
	}
	return true, nil
}
