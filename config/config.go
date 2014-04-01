package config

import (
	"code.google.com/p/gcfg"
)

type Config struct {
	UDP struct {
		Port int
		Endpoint string
	}
}

func ReadConfig(file string, config *Config) error {
	return gcfg.ReadFileInto(config, file)
}
