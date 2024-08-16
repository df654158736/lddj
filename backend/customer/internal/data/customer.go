package data

import (
	"context"
	"customer/internal/biz"
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

func (repo CustomerRepo) SetVerifyCode(phone_number string) (string, error) {
	code := "1234"
	status := repo.data.Rdb.Set(context.Background(), "CVC:"+phone_number, code, 60*time.Second)
	if _, err := status.Result(); err != nil {
		return "", err
	}
	return code, nil
}

func (repo *CustomerRepo) GetVerifyCode(phoneNumber string) string {
	// 实现获取验证码的逻辑
	code := repo.data.Rdb.Get(context.Background(), "CVC:"+phoneNumber)
	return code.Val()
}

func (repo *CustomerRepo) GetCustomerByTelephone(phoneNumber string) (*biz.Customer, error) {
	// 实现根据手机号获取客户信息的逻辑
	customer := &biz.Customer{}

	result := repo.data.Mdb.Where("telephone = ?", phoneNumber).First(customer)
	if result.Error == nil && result.RowsAffected > 0 {
		return customer, nil
	}

	return nil, result.Error
}
