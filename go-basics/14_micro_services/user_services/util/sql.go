package util

import (
	"log"
	"os"
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

func DBConnect(dbName string) (*gorm.DB, error) {
	/*
		注意：
			想要正确的处理 time.Time ，您需要带上 parseTime 参数，
			要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
			设置时间的位置 Local 为本地系统
	*/
	dsn := "root:xL1210...@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		// NameStrategy 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "svc_", // 表前缀 user => svc_users
			SingularTable: true,   // svc_user
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

// Paginate 分页
func Paginate(page, pageSize int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		// Limit 指定要检索的最大记录数
		// Offset 指定在开始返回记录之前要跳过的记录数
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
