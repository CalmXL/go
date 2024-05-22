package middwares

import (
	"13_gin/6_router-group/controller"
	"github.com/gin-gonic/gin"
)

var origins = []string{
	"http://localhost:5173",
	"http://localhost:3000",
}

func Router(r *gin.Engine) gin.HandlerFunc {
	phone := controller.Phone{}
	user := controller.User{}
	userGroup := r.Group("/user")
	phoneGroup := r.Group("/phone")

	return func(c *gin.Context) {
		phoneGroup.GET("/list", phone.GetList)
		phoneGroup.GET("/:id", phone.GetPhone)

		userGroup.GET("/list", user.GetList)
		userGroup.GET("/:id", user.GetUser)
	}
}
