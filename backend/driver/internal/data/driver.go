package data

import (
	"context"
	"time"

	"driver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type driverRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewDriverRepo(data *Data, logger log.Logger) biz.DriverRepo {
	return &driverRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *driverRepo) SaveToRedis(ctx context.Context, key string, content string, second int64) error {
	expiration := time.Duration(second) * time.Second
	status := r.data.Rdb.Set(context.Background(), "CVC:"+key, content, expiration)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}
