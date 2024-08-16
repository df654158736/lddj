package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type CustomerRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewCustomerRepo(data *Data, logger log.Logger) *CustomerRepo {
	return &CustomerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (cr CustomerRepo) SetVerifyCode(code string) error {
	status := cr.data.Rdb.Set(context.Background(), "CVC:"+code, "VerifyCode", 60*time.Second)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}
