package service

import (
	"context"
	"log"
	"math/rand"

	pb "verifyCode/api/s/v1"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) CreateVerifyCode(ctx context.Context, req *pb.CreateVerifyCodeRequest) (*pb.CreateVerifyCodeReply, error) {
	return &pb.CreateVerifyCodeReply{}, nil
}
func (s *VerifyCodeService) UpdateVerifyCode(ctx context.Context, req *pb.UpdateVerifyCodeRequest) (*pb.UpdateVerifyCodeReply, error) {
	return &pb.UpdateVerifyCodeReply{}, nil
}
func (s *VerifyCodeService) DeleteVerifyCode(ctx context.Context, req *pb.DeleteVerifyCodeRequest) (*pb.DeleteVerifyCodeReply, error) {
	return &pb.DeleteVerifyCodeReply{}, nil
}
func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	code := mRandomCode(int(req.Length), req.Type)
	log.Println("GetVerifyCode:", code)
	return &pb.GetVerifyCodeReply{Message: code}, nil
}
func (s *VerifyCodeService) ListVerifyCode(ctx context.Context, req *pb.ListVerifyCodeRequest) (*pb.ListVerifyCodeReply, error) {
	return &pb.ListVerifyCodeReply{}, nil
}

func mRandomCode(length int, tpye pb.Type) string {
	// TODO
	switch tpye {
	case pb.Type_DEFAULT:
		return ""
	case pb.Type_DIGIT:
		return mGetCode("0123456789", length)
	case pb.Type_LETTER:
		return mGetCode("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
	case pb.Type_MIXED:
		return mGetCode("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
	}
	return ""
}

func mGetCode(s string, length int) string {
	charsLen := len(s)

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		randIndex := rand.Intn(charsLen)
		result[i] = s[randIndex]
	}
	return string(result)
}
