package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type GinConfig struct {
	Ip   string `mapstructure:"ip"`
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
}

type Config struct {
	*GinConfig `mapstructure:"gin"`
	*DBConfig  `mapstructure:"db"`
}

var cfg = new(Config)

func Initialize() (*Config, error) {
	viper.SetConfigFile("./config/config.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("解析配置文件失败: %s", err.Error())
		}
	})

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
