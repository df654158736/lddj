package service

import (
	"context"
	pb "customer/api/customer/v1"
	"customer/internal/data"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	customerRepo *data.CustomerRepo
}

func NewCustomerService(cr *data.CustomerRepo) *CustomerService {
	return &CustomerService{
		customerRepo: cr,
	}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	println("GetVerifyCode")
	s.customerRepo.SetVerifyCode(req.PhoneNumber)
	return &pb.GetVerifyCodeReply{
		VerifyCode: "VerifyCode",
		Code:       200,
		Message:    "Message",
		ExpireTime: 60}, nil
}
