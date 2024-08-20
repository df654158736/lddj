package biz

import (
	"context"
	"fmt"
	"strconv"
	"time"
	map_servie "valuation/api/map"
	util "valuation/internal/util"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	GetRules(context.Context, uint, time.Time) (*PriceRules, error)
}

// GreeterUsecase is a Greeter usecase.
type ValuationUsecase struct {
	repo ValuationRepo
	log  *log.Helper
}

type PriceRules struct {
	gorm.Model
	PriceRuleWork
}

type PriceRuleWork struct {
	CityID      uint      `json:"city_id"`
	StartFee    int64     `json:"start_fee"`
	DistanceFee int64     `json:"distance_fee"`
	DurationFee int64     `json:"duration_fee"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}

// NewGreeterUsecase new a Greeter usecase.
func NewValuationUsecase(repo ValuationRepo, logger log.Logger) *ValuationUsecase {
	return &ValuationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ValuationUsecase) GetPrice(ctx context.Context, distance, duration string) (int32, error) {
	fmt.Println("distance:", distance, "---", "duration:", duration)
	rules, err := uc.repo.GetRules(ctx, 1, time.Now())
	if err != nil {
		return 0, err
	}
	distanceInt64, err := strconv.ParseInt(distance, 10, 64)
	if err != nil {
		return 0, err
	}

	var distanceStart int64 = 5
	total := rules.StartFee + rules.DistanceFee*(distanceInt64-distanceStart) + rules.DurationFee*distanceInt64

	return int32(total), nil
}

func (uc *ValuationUsecase) GetDrivingInfo(ctx context.Context, origin, destination string) (string, string, error) {
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

	distance, duration := result.Distance, result.Duration

	return distance, duration, nil
}
