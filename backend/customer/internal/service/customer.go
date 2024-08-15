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

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerReply, error) {
	return &pb.CreateCustomerReply{}, nil
}
func (s *CustomerService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerReply, error) {
	return &pb.UpdateCustomerReply{}, nil
}
func (s *CustomerService) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerReply, error) {
	return &pb.DeleteCustomerReply{}, nil
}
func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerReply, error) {
	return &pb.GetCustomerReply{}, nil
}
func (s *CustomerService) ListCustomer(ctx context.Context, req *pb.ListCustomerRequest) (*pb.ListCustomerReply, error) {
	return &pb.ListCustomerReply{}, nil
}
func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	println("GetVerifyCode")
	return &pb.GetVerifyCodeReply{}, nil
}
