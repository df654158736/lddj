package biz

import (
	"context"
	"database/sql"
	verifycode "driver/api/verify_code/v1"
	"driver/internal/util"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

const DriverStatusOut = "out"
const DriverStatusIn = "in"
const DriverStatusListen = "listen"
const DriverStatusStop = "stop"

// Greeter is a Greeter model.
type DriverModel struct {
	gorm.Model
	DriverWork
}

type DriverWork struct {
	Telephone     string         `gorm:"column:phone_number;type:varchar(20);not null;unique" json:"phone_number"`
	Token         sql.NullString `gorm:"column:token;type:varchar(2047);" json:"token"`
	Name          sql.NullString `gorm:"column:name;type:varchar(255);" json:"name"`
	Status        sql.NullString `gorm:"column:status;type:enum('out','in','listen','stop');" json:"status"`
	IdNumber      sql.NullString `gorm:"column:id_number;type:varchar(20);uniqueIndex;" json:"id_number"`
	IdImageA      sql.NullString `gorm:"column:id_image_a;type:varchar(255);" json:"id_image_a"`
	LicenseImageA sql.NullString `gorm:"column:license_image_a;type:varchar(255);" json:"license_image_a"`
	LicenseImageB sql.NullString `gorm:"column:license_image_b;type:varchar(255);" json:"license_image_b"`
	DistinctCode  sql.NullString `gorm:"column:distinct_code;type:varchar(16);" json:"distinct_code"`
	AuditAt       sql.NullTime   `gorm:"column:audit_at;type:datetime" json:"audit_at"`
	TelephoneBak  sql.NullString `gorm:"column:phone_code;type:varchar(16);" json:"telephone_bak"`
}

// GreeterRepo is a Greater repo.
type DriverRepo interface {
	SaveToRedis(context.Context, string, string, int64) error
	GetFromRedis(string) (string, error)
	SaveToMysql(string, string) (*DriverModel, error)
	GetModelByTelephone(string) (*DriverModel, error)
	UpdateTokenToMysql(string, string) (*DriverModel, error)
	GetTokenById(string) (string, error)
}

// GreeterUsecase is a Greeter usecase.
type DriverUsecase struct {
	repo DriverRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewDriverUsecase(repo DriverRepo, logger log.Logger) *DriverUsecase {
	return &DriverUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DriverUsecase) GetToken(id string) (string, error) {

	token, err := uc.repo.GetTokenById(id)

	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *DriverUsecase) SaveToken(phone string, token string) (string, error) {

	redisErr := uc.repo.SaveToRedis(context.Background(), "JWT:"+phone, token, 60*5)
	if redisErr != nil {
		return "", redisErr
	}

	_, mysqlErr := uc.repo.SaveToMysql(phone, token)

	if mysqlErr != nil {
		return "", mysqlErr
	}
	return token, nil
}

func (uc *DriverUsecase) Logout(phone string, token string) (string, error) {

	_, mysqlErr := uc.repo.UpdateTokenToMysql(phone, token)

	if mysqlErr != nil {
		return "", mysqlErr
	}
	return token, nil
}

func (uc *DriverUsecase) InitDriverInfo(cxt context.Context, phone_number string) (*DriverModel, error) {

	model, err := uc.repo.SaveToMysql(phone_number, "")

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (uc *DriverUsecase) GetModelByTelephone(phoneNumber string) (*DriverModel, error) {
	model, err := uc.repo.GetModelByTelephone(phoneNumber)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (uc *DriverUsecase) GetVerifyCodeFromRedis(phoneNumber string) string {
	// 实现获取验证码的逻辑
	code, err := uc.repo.GetFromRedis("CVC:" + phoneNumber)
	if err != nil {
		return ""
	}
	return code
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *DriverUsecase) GetVerifyCode(phone_number string, lenght int32, mtype int32) (string, error) {
	conn, err := util.GetGRPCConn("discovery:///VerifyCodeService")
	if err != nil {
		return "", err
	}

	defer func() {
		conn.Close()
	}()

	client := verifycode.NewVerifyCodeClient(conn)
	result, err := client.GetVerifyCode(context.Background(), &verifycode.GetVerifyCodeRequest{Length: lenght, Type: verifycode.Type(mtype)})

	if err != nil {
		println("2:", err)
		return "", err
	}
	fmt.Println("GRPC:", result)
	code := result.Message

	redisErr := uc.repo.SaveToRedis(context.Background(), "CVC:"+phone_number, code, 60*5)
	if redisErr != nil {
		println("2:", redisErr)
		return "", redisErr
	}
	return code, nil
}
