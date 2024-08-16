package data

import (
	"customer/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	redis "github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCustomerRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data := &Data{}

	cleanup := func() {
		//清理redis连接
		data.Rdb.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	//连接redis
	redisURL := fmt.Sprintf("redis://%s/1?dial_timeout=3s", c.Redis.Addr)
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		data.Rdb = nil
	}
	data.Rdb = redis.NewClient(options)

	return data, cleanup, nil
}
