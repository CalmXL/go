package main

import (
	"13_gin/3_Design/controller"
	"13_gin/3_Design/middware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middware.Cors())
	info := controller.Info{}

	r.GET("/info", info.GetInfo)
	r.GET("/info/:id", info.GetInfo2)
	r.POST("/create", info.Create)
	r.PUT("/update", info.Update)
	r.PATCH("/update/age", info.UpdateAge)
	r.DELETE("/delete", info.Delete)
	r.HEAD("/head", info.Head)
	r.OPTIONS("/methods", info.GetMethods)

	r.Run(":3000")
}
