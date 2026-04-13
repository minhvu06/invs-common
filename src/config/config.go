package config

import (
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Mode string `yaml:"mode"`
	} `yaml:"app"`

	JWT struct {
		Secret string `yaml:"secret"`
		Expire int64  `yaml:"expire"`
	} `yaml:"jwt"`

	Database struct {
		UserDB struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
		} `yaml:"user_db"`
		OrderDB struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
		} `yaml:"order_db"`
	} `yaml:"database"`

	RPC struct {
		UserRPC  string `yaml:"user_rpc"`
		OrderRPC string `yaml:"order_rpc"`
	} `yaml:"rpc"`
}

func MustLoadConfig(path string) *Config {
	var c Config
	conf.MustLoad(path, &c)
	return &c
}
