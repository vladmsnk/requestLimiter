package main

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"requestLimiter/pb"
)

func RunTestRequest(num int, client pb.InfoClient) {
	userID := uuid.NewString()
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("user_id", userID))
	for i := 0; i < num; i++ {
		info, err := client.GetInfo(ctx, &pb.GetInfoRequest{UserId: userID})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("Request #%d: %s", i, info.Info))
		}
	}
}

func main() {

	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewInfoClient(conn)
	RunTestRequest(10, client)
}
