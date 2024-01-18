package redis

import (
	"time"

	redislib "github.com/opensourceways/redis-lib"
)

const (
	defaultAddress  = "192.1.1.1"
	defaultPassword = "password"
	defaultDB       = 4
	defaultTimeOut  = 100 * time.Second
	defaultDBCert   = "xxxxx"
)

type Config struct {
	redislib.Config
}

func SetDefault() Config {
	return Config{
		redislib.Config{
			Address:  defaultAddress,
			Password: defaultPassword,
			DB:       defaultDB,
			DBCert:   defaultDBCert,
		},
	}
}
