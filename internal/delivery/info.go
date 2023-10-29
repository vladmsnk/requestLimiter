package delivery

import (
	"context"
	"requestLimiter/internal/useCase"
	"requestLimiter/pb"
)

type Delivery struct {
	useCase useCase.Info
	pb.UnimplementedInfoServer
}

func NewDelivery(uc useCase.Info) *Delivery {
	return &Delivery{useCase: uc}
}

func (d *Delivery) GetInfo(_ context.Context, in *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	userID := in.UserId
	info := d.useCase.GetInfo(userID)
	return &pb.GetInfoResponse{Info: info}, nil
}
