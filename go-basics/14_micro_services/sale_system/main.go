package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sale_system/config"
	"sale_system/middleware"
	"sale_system/util"
)

func main() {
	cfg, err := config.Initialize()

	if err != nil {
		zap.S().Fatalf("Failed to initialize config: %v", err.Error())
	}

	util.NewLogger(cfg.LogConfig)

	r := gin.Default()

	db, err := util.DBConnect(cfg.DBConfig)

	if err != nil {
		zap.S().Fatalf("Failed to connect to database: %v", err.Error())
	}

	r.Use(middleware.Cors(*cfg.AllowOrigins))

	if err := r.Run(cfg.GinConfig.IP + ":" + cfg.GinConfig.Port); err != nil {
		zap.S().Fatalf("Failed to run gin service: %v", err.Error())
		return
	}
}
