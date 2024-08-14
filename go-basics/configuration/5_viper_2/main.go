package main

import (
	"configuration/5_viper_2/config"
	"configuration/5_viper_2/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Product struct {
	Id    int32
	Name  string
	Price float32
}

func main() {
	cfg, err := config.Initialize()

	if err != nil {
		log.Fatalf("初始化失败：%s", err.Error())
		return
	}

	db, err := util.DBConnect(cfg.DBConfig)

	if err != nil {
		log.Fatalf("数据库连接失败: %s", err.Error())
		return
	}

	r := gin.Default()

	r.GET("/list", func(c *gin.Context) {
		var productList []Product

		result := db.Model(&productList).Find(&productList)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 1001,
				"msg":  "查询数据库失败",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": productList,
		})
	})

	if err := r.Run(fmt.Sprintf("%s:%s", cfg.GinConfig.Ip, cfg.GinConfig.Port)); err != nil {
		log.Fatalf("服务器程序运行失败: %s", err.Error())
		return
	}
}
