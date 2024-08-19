package biz

import (
	"context"
	"fmt"
	map_servie "valuation/api/map"
	util "valuation/internal/util"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
)

// GreeterRepo is a Greater repo.
type ValuationRepo interface {
	Save(context.Context, *ValuationUsecase) (*ValuationUsecase, error)
	Update(context.Context, *ValuationUsecase) (*ValuationUsecase, error)
	FindByID(context.Context, int64) (*ValuationUsecase, error)
	ListByHello(context.Context, string) ([]*ValuationUsecase, error)
	ListAll(context.Context) ([]*ValuationUsecase, error)
}

// GreeterUsecase is a Greeter usecase.
type ValuationUsecase struct {
	repo ValuationRepo
	log  *log.Helper
}

type DirectionDrivingResp struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	Count  string `json:"count"`
	Route  struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		TaxiCost    string `json:"taxi_cost"`
		Paths       []struct {
			Distance string `json:"distance"`
			Duration string `json:"duration"`
		}
	} `json:"route"`
}

// NewGreeterUsecase new a Greeter usecase.
func NewValuationUsecase(repo ValuationRepo, logger log.Logger) *ValuationUsecase {
	return &ValuationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (vubiz *ValuationUsecase) GetDrivingInfo(ctx context.Context, origin, destination string) (string, string, error) {
	conn, err := util.GetGRPCConn("MapService")
	if err != nil {
		return "", "", err
	}

	defer func() {
		conn.Close()
	}()

	client := map_servie.NewMapClient(conn)
	fmt.Println("origin:", origin, "destination:", destination)
	result, err := client.GetDrivingInfo(context.Background(), &map_servie.GetDrivingInfoRequest{Origin: origin, Destination: destination})

	if err != nil {
		println("2:", err)
		return "", "", err
	}
	fmt.Println(result)

	distance, duration := result.Destination, result.Duration

	return distance, duration, nil
}
