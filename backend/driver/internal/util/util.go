package util

import (
	"context"
	"log"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"

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
