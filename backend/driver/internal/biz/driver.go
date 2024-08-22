package biz

import (
	"context"
	verifycode "driver/api/verify_code/v1"
	"driver/internal/util"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
)

// Greeter is a Greeter model.
type DriverModel struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type DriverRepo interface {
	SaveToRedis(context.Context, string, string, int64) error
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

	redisErr := uc.repo.SaveToRedis(context.Background(), phone_number, code, 60*5)
	if redisErr != nil {
		println("2:", redisErr)
		return "", redisErr
	}
	return code, nil
}
