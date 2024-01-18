package redis

import (
	redislib "github.com/opensourceways/redis-lib"
)

var RedisCli = redislib.DAO()

func Init(cfg *Config) {
	redislib.Init(&cfg.Config, false)
}

func Close() error {
	return redislib.Close()
}
