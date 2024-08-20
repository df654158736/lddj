package data

import (
	"context"
	"fmt"
	"time"

	"valuation/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type valuationRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewValuationRepo(data *Data, logger log.Logger) biz.ValuationRepo {
	return &valuationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *valuationRepo) Save(ctx context.Context, g *biz.ValuationUsecase) (*biz.ValuationUsecase, error) {
	return g, nil
}

func (r *valuationRepo) Update(ctx context.Context, g *biz.ValuationUsecase) (*biz.ValuationUsecase, error) {
	return g, nil
}

func (r *valuationRepo) FindByID(context.Context, int64) (*biz.ValuationUsecase, error) {
	return nil, nil
}

func (r *valuationRepo) ListByHello(context.Context, string) ([]*biz.ValuationUsecase, error) {
	return nil, nil
}

func (r *valuationRepo) ListAll(context.Context) ([]*biz.ValuationUsecase, error) {
	return nil, nil
}

func (r *valuationRepo) GetRules(ctx context.Context, city_id uint, current_time time.Time) (*biz.PriceRules, error) {
	p := &biz.PriceRules{}
	// result := r.data.Mdb.Where("city_id=? and start_at >= ? and end_at <= ?", city_id, current_time, current_time).First(p)
	result := r.data.Mdb.Where("city_id=?", city_id).First(p)
	if result.Error != nil {
		println("************", result.Error)
		return nil, result.Error
	}
	fmt.Println(p)
	return p, nil
}
