package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"log"
	"sale_system/util"
)

type GinConfig struct {
	IP   string `mapstructure:"ip"`
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Prefix   string `mapstructure:"prefix"`
}

type LogConfig struct {
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

type AllowOrigins = []string

type Config struct {
	*GinConfig    `mapstructure:"gin"`
	*DBConfig     `mapstructure:"db"`
	*LogConfig    `mapstructure:"log"`
	*AllowOrigins `mapstructure:"allow_origins"`
}

var cfg = new(Config)

var FileConfig = map[string]string{
	"PREFIX": "config",
	"ENV":    "SALE_SYSTEM_ENV",
}

var ZapLevel = zapcore.InfoLevel

func getEnvFile() string {
	var configFile string
	isDev := util.GetEnv(FileConfig["ENV"])
	fmt.Println("isDEV: ", isDev)
	if isDev {
		configFile = fmt.Sprintf("./config/%s-dev.yaml", FileConfig["PREFIX"])
	} else {
		configFile = fmt.Sprintf("./config/%s-prod.yaml", FileConfig["PREFIX"])
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
