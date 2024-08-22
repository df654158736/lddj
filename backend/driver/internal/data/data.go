package data

import (
	"driver/internal/biz"
	"driver/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDriverRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
	Mdb *gorm.DB
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
	redisURL := fmt.Sprintf("redis://%s/2?dial_timeout=3s", c.Redis.Addr)
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		data.Rdb = nil
		log.Fatal(err)
	}
	data.Rdb = redis.NewClient(options)

	//连接mysql
	dsn := c.Database.Source
	data.Mdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		data.Mdb = nil
		log.Fatal(err)
	}

	//初始化表结构
	migrateTables(data.Mdb)

	return data, cleanup, nil
}

func migrateTables(db *gorm.DB) {
	if err := db.AutoMigrate(&biz.DriverModel{}); err != nil {
		log.Fatal(err)
	}
}
