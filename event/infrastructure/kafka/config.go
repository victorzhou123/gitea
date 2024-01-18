package kafka

import (
	kfklib "github.com/opensourceways/kafka-lib/agent"
)

const (
	defaultAddress  = "192.1.1.1"
	defaultVersion  = "2.1.0"
	defaultMQCert   = "xxxx"
	defaultUserName = "xxxxx"
	defaultPassword = "xxxxx"
)

type Config struct {
	kfklib.Config
}

func SetDefault() Config {
	return Config{
		kfklib.Config{
			Address:        defaultAddress,
			Version:        defaultVersion,
			MQCert:         defaultMQCert,
			Username:       defaultUserName,
			Password:       defaultPassword,
			SkipCertVerify: true,
		},
	}
}
