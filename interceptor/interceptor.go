package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"requestLimiter/limiter"
	"time"
)

func LimiterInterceptor(limiter limiter.Limiter) grpc.UnaryServerInterceptor {
	requestTime := time.Now()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		userID := ctx.Value("userID").(string)
		allowed, err := limiter.Allow(ctx, userID, requestTime)
		if err != nil {
			return nil, err
		}
		if !allowed {
			return nil, status.Error(codes.ResourceExhausted, "Too many requests")
		}

		return handler(ctx, req)
	}
}
