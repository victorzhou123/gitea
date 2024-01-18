package kafka

import (
	kfklib "github.com/opensourceways/kafka-lib/agent"
)

const (
	defaultAddress = "172.23.0.7:9092"
	defaultVersion = "2.1.0"
)

type Config struct {
	kfklib.Config
}

func SetDefault() Config {
	return Config{
		kfklib.Config{
			Address:        defaultAddress,
			Version:        defaultVersion,
			SkipCertVerify: true,
		},
	}
}
