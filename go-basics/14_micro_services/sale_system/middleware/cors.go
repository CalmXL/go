package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sale_system/config"
)

func Cors(origin config.AllowOrigins) gin.HandlerFunc {
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = origin
	corsCfg.AllowOrigins = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsCfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsCfg.AllowCredentials = true

	return cors.New(corsCfg)
}
