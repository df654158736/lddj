package service

import (
	"context"
	pb "customer/api/customer/v1"
	"customer/internal/data"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	CustomerRepo *data.CustomerRepo
}

func NewCustomerService(cr *data.CustomerRepo) *CustomerService {
	return &CustomerService{
		CustomerRepo: cr,
	}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	println("GetVerifyCode")
	code, err := s.CustomerRepo.SetVerifyCode(req.PhoneNumber)
	if err != nil {
		return &pb.GetVerifyCodeReply{
			Code:       400,
			Message:    err.Error(),
			VerifyCode: code,
			ExpireTime: 60}, nil
	}
	return &pb.GetVerifyCodeReply{
		Code:       200,
		Message:    "Message",
		VerifyCode: code,
		ExpireTime: 60}, nil
}

func (s *CustomerService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	fmtStr := fmt.Sprintf("Login: %s, %s", req.PhoneNumber, req.VerifyCode)
	println(fmtStr)

	// TODO: verify code
	code := s.CustomerRepo.GetVerifyCode(req.PhoneNumber)
	codeStr := fmt.Sprintf("code: %s", code)
	println(codeStr)

	if code != req.VerifyCode {
		return &pb.LoginReply{
			Code:       400,
			Message:    "Verify code error",
			ExpireTime: 60}, nil
	}

	// TODO: 用户是否存在
	customer, err := s.CustomerRepo.GetCustomerByTelephone(req.PhoneNumber)
	if err != nil {
		return &pb.LoginReply{
			Code:       400,
			Message:    err.Error(),
			ExpireTime: 60}, nil
	}

	// TODO: 生成token
	token, err := s.CustomerRepo.GenerateTokenAndSave(customer)
	if err != nil {
		return &pb.LoginReply{
			Code:       400,
			Message:    err.Error(),
			ExpireTime: 60}, nil
	}

	return &pb.LoginReply{
		Code:       200,
		Message:    "Message",
		Token:      token,
		ExpireTime: 60}, nil
}

func (s *CustomerService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {

	claims, _ := jwt.FromContext(ctx)

	//断言
	claimsMap := claims.(jwtv5.MapClaims)
	id := claimsMap["jti"].(string)

	log.Info("Logout:", "customerId:", id)

	err := s.CustomerRepo.DeleteTokenById(id)
	if err != nil {
		return &pb.LogoutReply{
			Code:    400,
			Message: err.Error(),
		}, nil
	} else {
		return &pb.LogoutReply{
			Code:    200,
			Message: "Message",
		}, nil
	}
}
