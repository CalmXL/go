package controller

import (
	"13_gin/6_router-group/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
}

func (u *User) GetList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data.Users,
	})
}

func (u *User) GetUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	fmt.Println(id)
	for _, user := range data.Users {
		if user.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": user,
			})
			break
		}
	}
}
