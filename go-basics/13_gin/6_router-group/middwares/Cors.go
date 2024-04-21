package middwares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("origin")

		for _, o := range origins {
			if o == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", "GET,POST")
				c.Header("Access-Control-Allow-Headers", "Content-Type")

				if c.Request.Method == "OPTIONS" {
					c.JSON(http.StatusOK, nil)
					c.Abort()
					return
				}

				c.Next()
			}
		}

	}
}
