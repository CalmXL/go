package main

import (
	"13_gin/6_router-group/middwares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middwares.Cors())
	r.Use(middwares.Router(r))

	if err := r.Run(":3000"); err != nil {
		fmt.Println("服务器启动失败!")
	}
}
