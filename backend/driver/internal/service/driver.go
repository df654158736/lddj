package service

import (
	"context"

	pb "driver/api/driver/v1"
	"driver/internal/biz"
)

type DriverService struct {
	pb.UnimplementedDriverServer
	uc *biz.DriverUsecase
}

func NewDriverService(vuc *biz.DriverUsecase) *DriverService {
	return &DriverService{
		uc: vuc,
	}
}

func (s *DriverService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	code, err := s.uc.GetVerifyCode("123456", req.Length, int32(req.Type))
	if err != nil {
		return &pb.GetVerifyCodeReply{Message: err.Error()}, nil
	}
	return &pb.GetVerifyCodeReply{Message: code}, nil
}
