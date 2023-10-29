package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"requestLimiter/limiter"
)

func LimiterInterceptor(limiter limiter.Limiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		ok, err := limiter.Limit(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !ok {
			return nil, status.Error(codes.ResourceExhausted, "Too many requests")
		}

		return handler(ctx, req)
	}
}
