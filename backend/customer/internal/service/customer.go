package service

import (
	"context"

	pb "customer/api/customer/v1"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	println("GetVerifyCode")
	return &pb.GetVerifyCodeReply{}, nil
}
