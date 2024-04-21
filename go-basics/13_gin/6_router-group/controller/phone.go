package controller

import (
	"13_gin/6_router-group/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Phone struct {
}

func (p *Phone) GetList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data.Phones,
	})
}

func (p *Phone) GetPhone(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	for _, phone := range data.Phones {
		if phone.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": phone,
			})
			break
		}
	}
}
