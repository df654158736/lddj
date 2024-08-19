package data

import (
	"context"
	verifycode "customer/api/verify_code"
	"customer/internal/biz"
	util "customer/internal/util"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

func (repo CustomerRepo) GetVerifyCodeGRPC(phone_number string, lenght int32, mtype int32) (string, error) {
	conn, err := util.GetGRPCConn()
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
	fmt.Println(result)
	code := result.Message
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

func (repo *CustomerRepo) GetToken(id string) (string, error) {
	customer := &biz.Customer{}
	result := repo.data.Mdb.Where("Id = ?", id).First(customer)
	if result.Error == nil && result.RowsAffected > 0 {
		return customer.Token, nil
	}

	return "", result.Error
}

func (repo *CustomerRepo) DeleteTokenById(id string) error {
	customer := &biz.Customer{}
	result := repo.data.Mdb.Where("Id = ?", id).First(customer)
	if result.Error == nil && result.RowsAffected > 0 {

		customer.Token = ""
		customer.TokenCreatedAt = sql.NullTime{Time: time.Time{}, Valid: false}
		updateResult := repo.data.Mdb.Save(customer)
		if updateResult.Error == nil && updateResult.RowsAffected > 0 {
			return nil
		} else {
			return updateResult.Error
		}
	}

	return result.Error
}

func (repo *CustomerRepo) GetCustomerByTelephone(phoneNumber string) (*biz.Customer, error) {
	// 实现根据手机号获取客户信息的逻辑
	customer := &biz.Customer{}

	result := repo.data.Mdb.Where("telephone = ?", phoneNumber).First(customer)
	if result.Error == nil && result.RowsAffected > 0 {
		return customer, nil
	}

	//如何customer没有找到，则创建一个新的customer
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		customer.Telephone = phoneNumber
		createResult := repo.data.Mdb.Create(customer)
		if createResult.Error == nil && createResult.RowsAffected > 0 {
			return customer, nil
		}
		return nil, createResult.Error
	}

	return nil, result.Error
}

func (repo *CustomerRepo) GenerateTokenAndSave(customer *biz.Customer) (string, error) {
	duration := 24 * time.Hour
	log.Info(repo.log, "duration", duration)
	secret := "secret"
	log.Info(repo.log, "secret", secret)

	//获取jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		//签发方，签发机构
		Issuer: "DD",
		//说明
		Subject: "customer verification",
		//签发给谁使用
		Audience: []string{"customer"},
		//签发时间
		IssuedAt: jwt.NewNumericDate(time.Now()),
		//何时启用
		NotBefore: jwt.NewNumericDate(time.Now()),
		//ID 标识
		ID: fmt.Sprintf("%d", customer.ID),
		//有效期至
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	log.Info(repo.log, "token", tokenString)

	//存储 jwt token
	status := repo.data.Rdb.Set(context.Background(), "JWT:"+customer.Telephone, tokenString, duration)
	if _, err := status.Result(); err != nil {
		return "", err
	}

	//存储到mysql
	customer.Token = tokenString
	customer.TokenCreatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	sqlResult := repo.data.Mdb.Save(customer)
	if sqlResult.Error != nil {
		return "", sqlResult.Error
	}

	return tokenString, nil
}
func (repo *CustomerRepo) VerifyToken(tokenString string) (*biz.Customer, error) {
	// 实现验证token的逻辑
	return nil, nil
}
