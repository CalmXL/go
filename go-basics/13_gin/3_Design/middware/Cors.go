package middware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATH,OPTIONS,HEAD")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
