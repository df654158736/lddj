package util

import (
	"context"
	"log"
	"time"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/golang-jwt/jwt/v5"

	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

func GetGRPCConn(endpoint string) (*grpc.ClientConn, error) {
	//1，获取consul客户端
	consulConfig := capi.DefaultConfig()

	//通过配置文件拿到consul服务的地址
	consulConfig.Address = "192.168.86.133:8500"
	consulClient, err := capi.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}

	//2，获取服务发现管理器
	discovery := consul.New(consulClient)

	//连接grpc服务
	// endpoint := "discovery:///VerifyCodeService"
	conn, err := kgrpc.DialInsecure(context.Background(), kgrpc.WithEndpoint(endpoint), kgrpc.WithDiscovery(discovery))
	if err != nil {
		println("GetGRPCConn err:", err)
		return nil, err
	}

	return conn, nil
}

func GenerateJWTToken(user_id string) (string, error) {
	duration := 24 * time.Hour
	secret := "secret"

	println("GenerateJWTTokenById：", user_id)

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
		ID: user_id,
		//有效期至
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
