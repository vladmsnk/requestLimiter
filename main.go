package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"requestLimiter/client"
	"requestLimiter/config"
	"requestLimiter/interceptor"
	"requestLimiter/internal/delivery"
	"requestLimiter/internal/repository"
	"requestLimiter/internal/useCase/info"
	"requestLimiter/limiter"
	"requestLimiter/pb"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rClient := client.NewRedisClient(cfg)

	lm := limiter.NewLimiter(5, 60, rClient)

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.LimiterInterceptor(lm)))

	repo := repository.NewRepository()
	uc := info.NewUseCase(repo)
	infoDelivery := delivery.NewDelivery(uc)

	pb.RegisterInfoServer(server, infoDelivery)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
