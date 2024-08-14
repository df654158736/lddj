package service

import (
	"context"
	"log"

	pb "verifyCode/api/a"
)

type DemoService struct {
	pb.UnimplementedDemoServer
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	log.Println("GetDemo")
	return &pb.GetDemoReply{Message: "ok"}, nil
}
