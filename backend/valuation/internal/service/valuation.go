package service

import (
	"context"
	"fmt"

	pb "valuation/api/valuation/v1"
	"valuation/internal/biz"
)

type ValuationService struct {
	pb.UnimplementedValuationServer
	uc *biz.ValuationUsecase
}

func NewValuationService(vubiz *biz.ValuationUsecase) *ValuationService {
	return &ValuationService{
		uc: vubiz,
	}
}

func (s *ValuationService) GetEstimatePrice(ctx context.Context, req *pb.EstimatePriceRequest) (*pb.EstimatePriceReply, error) {
	distance, duration, err := s.uc.GetDrivingInfo(ctx, req.Origin, req.Destination)
	if err != nil {
		fmt.Printf("GetDrivingInfo error: %v\n", err)
		return &pb.EstimatePriceReply{}, nil
	}
	price, err := s.uc.GetPrice(ctx, distance, duration)
	if err != nil {
		fmt.Printf("GetPrice error: %v\n", err)
		return &pb.EstimatePriceReply{}, nil
	}
	return &pb.EstimatePriceReply{
		Origin:      distance,
		Destination: duration,
		Price:       price,
		Code:        1,
		Message:     "",
	}, nil
}
