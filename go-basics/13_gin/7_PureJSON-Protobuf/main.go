package main

import (
	"13_gin/7_PureJSON-Protobuf/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*

	  Pure JSON
		{
			id: 1,
			name: 'title',
			html: '<h1>This is TITLE</h1>'
		}
*/
func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		c.PureJSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gin.H{
				"id":    1,
				"name":  "xulei",
				"title": "<h1>This is TITLE</h1>",
			},
		})
	})

	r.GET("/protobuf", func(c *gin.Context) {
		user := &proto.User{
			Uid:    1,
			Name:   "xiaoyesensen",
			Age:    38,
			Gender: "male",
			Languages: []string{
				"Java",
				"Go",
				"JavaScript",
			},
		}

		c.ProtoBuf(http.StatusOK, user)
	})

	r.Run()
}
