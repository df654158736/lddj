package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"driver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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

// GetToRedis implements biz.DriverRepo.
func (r *driverRepo) GetFromRedis(key string) (string, error) {
	// 实现获取验证码的逻辑
	code := r.data.Rdb.Get(context.Background(), key)
	return code.Val(), nil
}

// SaveToMysql implements biz.DriverRepo.
func (r *driverRepo) SaveToMysql(phone string, token string) (*biz.DriverModel, error) {
	model := &biz.DriverModel{}

	result := r.data.Mdb.Where("phone_number = ?", phone).First(model)
	if result.Error == nil && result.RowsAffected > 0 {
		// 如果token不为空，则更新token
		if token != "" {
			model.Token = sql.NullString{String: token, Valid: true}
			updateResult := r.data.Mdb.Save(model)
			if updateResult.Error == nil && updateResult.RowsAffected > 0 {
				return model, nil
			} else {
				return nil, updateResult.Error
			}
		}

		return model, nil
	}

	//如何customer没有找到，则创建一个新的customer
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		model.Telephone = phone
		model.Token = sql.NullString{String: token, Valid: true}
		model.Status = sql.NullString{String: "stop", Valid: true}
		createResult := r.data.Mdb.Create(model)
		if createResult.Error == nil && createResult.RowsAffected > 0 {
			return model, nil
		}
		return nil, createResult.Error
	}

	return nil, result.Error
}

func (repo *driverRepo) SaveToRedis(ctx context.Context, key string, content string, second int64) error {
	expiration := time.Duration(second) * time.Second
	status := repo.data.Rdb.Set(context.Background(), key, content, expiration)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}

func (repo *driverRepo) GetModelByTelephone(phoneNumber string) (*biz.DriverModel, error) {
	// 实现根据手机号获取客户信息的逻辑
	model := &biz.DriverModel{}

	result := repo.data.Mdb.Where("phone_number = ?", phoneNumber).First(model)
	if result.Error == nil && result.RowsAffected > 0 {
		return model, nil
	}

	//如何customer没有找到，则创建一个新的customer
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		model.Telephone = phoneNumber
		createResult := repo.data.Mdb.Create(model)
		if createResult.Error == nil && createResult.RowsAffected > 0 {
			return model, nil
		}
		return nil, createResult.Error
	}

	return nil, result.Error
}

func (repo *driverRepo) GetTokenById(id string) (string, error) {
	model := &biz.DriverModel{}
	println("-------------id: ", id)
	result := repo.data.Mdb.Where("Id = ?", id).First(model)
	if result.Error == nil && result.RowsAffected > 0 {
		return model.Token.String, nil
	}

	return "", result.Error
}
