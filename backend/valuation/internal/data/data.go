package data

import (
	"time"
	"valuation/internal/biz"
	"valuation/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewValuationRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Mdb *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data := &Data{}
	//连接mysql
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		db = nil
		log.Fatal(err)
	}

	data.Mdb = db

	//初始化表结构
	migrateTables(data.Mdb)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}

func migrateTables(db *gorm.DB) {
	if err := db.AutoMigrate(&biz.PriceRules{}); err != nil {
		log.Fatal(err)
	}

	row := []biz.PriceRules{
		{
			Model: gorm.Model{ID: 1},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      1,
				StartFee:    8,
				DistanceFee: 35,
				DurationFee: 0,
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(time.Minute * 10),
			},
		}, {
			Model: gorm.Model{ID: 2},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      2,
				StartFee:    10,
				DistanceFee: 40,
				DurationFee: 0,
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(time.Minute * 10),
			},
		}, {
			Model: gorm.Model{ID: 3},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      3,
				StartFee:    12,
				DistanceFee: 45,
				DurationFee: 0,
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(time.Minute * 10),
			},
		},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&row)

}
