package gorani

import (
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/api/config"
	"github.com/sunho/gorani-reader/server/api/log"
)

type Gorani struct {
	Config config.Config
	Mysql  *gorm.DB
	Redis  *redis.Client
	Logger log.Logger
}

func NewGorani(conf config.Config) (*Gorani, error) {
	mysql, err := createMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	redis, err := createRedisConn(conf)
	if err != nil {
		return nil, err
	}

	l, err := createLogger(conf)
	if err != nil {
		return nil, err
	}

	gorn := &Gorani{
		Config: conf,
		Mysql:  mysql,
		Redis:  redis,
		Logger: l,
	}

	return gorn, nil
}

func createLogger(conf config.Config) (log.Logger, error) {
	switch conf.LoggerType {
	case config.LoggerTypeStdout:
		return log.NewStdoutLogger(), nil

	case config.LoggerTypeFluent:
		l, err := log.NewFluentLogger(conf.FluentHost, conf.FluentPort)
		if err != nil {
			return nil, err
		}

		return l, nil

	case config.LoggerTypeBoth:
		l, err := log.NewFluentLogger(conf.FluentHost, conf.FluentPort)
		if err != nil {
			return nil, err
		}
		s := log.NewStdoutLogger()

		return log.NewBothLogger(l, s), nil

	default:
		return nil, fmt.Errorf("Allowed logger types are stdout, fluent, both")
	}
}

func createMysqlConn(conf config.Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", conf.MysqlURL)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(conf.MysqlConnectionPoolSize)

	return db, nil
}

func createRedisConn(conf config.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(conf.RedisURL)
	if err != nil {
		return nil, err
	}

	opt.PoolSize = conf.RedisConnectionPoolSize

	client := redis.NewClient(opt)
	_, err = client.Ping().Result()

	return client, err
}