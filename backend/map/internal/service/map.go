package service

import (
	"context"

	pb "map/api/map/v1"
	"map/internal/biz"
)

type MapService struct {
	pb.UnimplementedMapServer
	uc *biz.MapUsecase
}

func NewMapService() *MapService {
	return &MapService{}
}

func (s *MapService) GetDrivingInfo(ctx context.Context, req *pb.GetDrivingInfoRequest) (*pb.GetDrivingInfoReply, error) {
	distance, duration, err := s.uc.GetDrivingInfo(req.Origin, req.Destination)
	if err != nil {
		return &pb.GetDrivingInfoReply{}, nil
	}
	return &pb.GetDrivingInfoReply{Origin: req.Origin, Destination: req.Destination, Distance: distance, Duration: duration}, nil
}
