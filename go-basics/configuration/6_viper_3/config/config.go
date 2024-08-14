package config

import (
	"fmt"
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
var fileConfig = map[string]string{
	"PREFIX": "config",
	"ENV":    "SALE_SYSTEM_ENV",
}

func getEnv(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func getEnvFile() string {
	var configFile string
	isDev := getEnv(fileConfig["ENV"])
	fmt.Println("isDEV: ", isDev)
	if isDev {
		configFile = fmt.Sprintf("./config/%s-dev.yaml", fileConfig["PREFIX"])
	} else {
		configFile = fmt.Sprintf("./config/%s-prod.yaml", fileConfig["PREFIX"])
	}

	return configFile
}

func Initialize() (*Config, error) {
	cfgFilePath := getEnvFile()
	fmt.Println(cfgFilePath)
	viper.SetConfigFile(cfgFilePath)

	// viper.SetConfigFile("./config/config.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("解析配置文件失败: %s", err.Error())
		}
	})

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(22)
		return nil, err
	}
	fmt.Println(72, cfg)
	return cfg, nil
}
