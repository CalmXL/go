package config

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
