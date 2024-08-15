package util

import (
	"fmt"
	"log"
	"os"
	"sale_system/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// gorm 的默认 logger 日志
var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer

	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  true,
	},
)

func DBConnect(cfg *config.DBConfig) (*gorm.DB, error) {
	/*
		注意：
			想要正确的处理 time.Time ，您需要带上 parseTime 参数，
			要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
			设置时间的位置 Local 为本地系统
	*/
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local)",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		// NameStrategy 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Prefix, // 表前缀 user => svc_users
			SingularTable: true,       // svc_user
		},
	})

	return db, err
}

func DBClose(db *gorm.DB) {
	sqlDB, errDB := db.DB()

	if errDB != nil {
		// 生成格式化的日志
		log.Fatalf("Failed to start server: %v", errDB)
		return
	}

	defer sqlDB.Close()
}
