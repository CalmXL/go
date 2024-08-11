package main

import (
	"configuration/3_go_config_style/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// db, err := DBConnect(
	// 	config.DBHost,
	// 	config.DBPort,
	// 	config.DBUser,
	// 	config.DBPassword,
	// 	config.DBName,
	// )

	// db, err := DBConnect(
	// 	config.DBConfig["HOST"],
	// 	config.DBConfig["PORT"],
	// 	config.DBConfig["USER"],
	// 	config.DBConfig["PASSWORD"],
	// 	config.DBConfig["NAME"],
	// )

	// db, err := DBConnect(
	// 	*config.DBHost,
	// 	*config.DBPort,
	// 	*config.DBUser,
	// 	*config.DBPassword,
	// 	*config.DBName,
	// )

	db, err := DBConnect(
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
	)

	if err != nil {
		log.Fatalf("数据库连接失败的原因: %s", err.Error())
		return
	}

	type Product struct {
		Id    int32
		Name  string
		Price float32
	}

	r.GET("/list", func(c *gin.Context) {
		var products []Product

		db.Model(&Product{}).Find(&products)
		c.JSON(http.StatusOK, products)
	})

	// cfg := fmt.Sprintf("%s:%s", config.GinHost, config.GinPort)
	// cfg := fmt.Sprintf("%s:%s", config.GINConfig["HOST"], config.GINConfig["PORT"])
	// cfg := fmt.Sprintf("%s:%s", *config.GinHost, *config.GinPort)
	cfg := fmt.Sprintf("%s:%s", config.Gin.Host, config.Gin.Port)

	if err := r.Run(cfg); err != nil {
		log.Fatalf("服务器程序启动失败: %s", err.Error())
		return
	}
}

func DBConnect(
	host string,
	port string,
	user string,
	password string,
	dbName string,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
