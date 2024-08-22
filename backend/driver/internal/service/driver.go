package service

import (
	"context"
	"fmt"
	"strconv"

	pb "driver/api/driver/v1"
	"driver/internal/biz"
	"driver/internal/util"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
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
func (s *DriverService) SubmitPhone(ctx context.Context, req *pb.SubmitPhoneRequest) (*pb.SubmitPhoneReply, error) {
	model, err := s.uc.InitDriverInfo(ctx, req.Phone)
	if err != nil {
		return &pb.SubmitPhoneReply{Message: "司机号码提交失败", Code: 1}, nil
	}
	return &pb.SubmitPhoneReply{Message: model.Status.String}, nil
}
func (s *DriverService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	fmtStr := fmt.Sprintf("Login: %s, %s", req.PhoneNumber, req.VerifyCode)
	println(fmtStr)
	// TODO: verify code
	code := s.uc.GetVerifyCodeFromRedis(req.PhoneNumber)
	codeStr := fmt.Sprintf("code: %s", code)
	println(codeStr)

	if code != req.VerifyCode {
		return &pb.LoginReply{
			Code:       400,
			Message:    "Verify code error",
			ExpireTime: 60}, nil
	}

	// TODO: 用户是否存在
	customer, err := s.uc.GetModelByTelephone(req.PhoneNumber)

	if err != nil {
		return &pb.LoginReply{
			Code:       400,
			Message:    err.Error(),
			ExpireTime: 60}, nil
	}

	println("customer:", customer.ID)
	idString := strconv.FormatUint(uint64(customer.ID), 10)
	// TODO: 生成token
	token, err := util.GenerateJWTToken(idString)
	if err != nil {
		return &pb.LoginReply{
			Code:       400,
			Message:    err.Error(),
			ExpireTime: 60}, nil
	}

	_, err = s.uc.SaveToken(req.PhoneNumber, token)
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

func (s *DriverService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	claims, _ := jwt.FromContext(ctx)

	//断言
	claimsMap := claims.(jwtv5.MapClaims)
	id := claimsMap["jti"].(string)

	log.Info("Logout:", "customerId:", id)

	_, err := s.uc.Logout(id, "-1")
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

func (s *DriverService) GetToken(id string) (string, error) {

	token, err := s.uc.GetToken(id)

	if err != nil {
		return "", err
	}
	return token, nil
}
