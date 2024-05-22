package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*

	  什么是复杂请求？

		在客户端发起请求，而实际客户端会自动的先发一次预请求，到服务器验证是否有权限发起实际请求
		如果有权限，客户端会发起实际的请求
		如果没有权限: 客户端会拦截服务器的相应

		预请求：是 OPTIOINS

		什么是简单请求？
		不是复杂请求的，都是简单请求

		1. 只会出现在 GET/POST/HEAD 请求方式中
		2. 请求头: Accept， Accept-Language, Content-Language, Content-Type ...
		3. Content-Type: application/x-www-form-urlencoded, text/plain, multiple/form-data

		PUT/PATCH/DELETE 复杂请求
	  POST Content-Type: application/json

		复杂请求:
			1. 客户端在数据请求之前多一次 OPTIONS 请求（预请求 - 权限验证请求）
			2. 服务端程序会多响应一次
			3. 预请求只会出现在不同源的异步请求中出现

		对服务器数据的更新删除创建的操作，对服务器数据安全有重大影响，所以要进行一次
		预检

		CSRF/XSRF Cross-Site request forgery 跨站请求伪造
		cors 方案: 在服务器设置 Access-Control-Allow-Origin

		XSS：Cross site Scripting 跨站脚本攻击


		Access to XMLHttpRequest at 'http://localhost:3000/user'
		from origin 'http://localhost:5173'
		has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header
		is present on the requested resource.

		5173 端口异步请求 :3000/user 接口因为同源策略被拦截了
		因为请求的资源没有出现 Access-Control-Allow-Origin 的头信息设置

		Access-Control-Allow-Origin
		访问    控制    允许   源
*/
func main() {
	r := gin.Default()

	// 除了 html, .tpl 也可以
	r.LoadHTMLGlob("4_request/*")

	origins := []string{
		"http://localhost:5173",
		"http://localhost:3000",
		"http://localhost:3300",
	}

	// 设置 cors
	r.Use(func(c *gin.Context) {

		origin := c.GetHeader("origin")
		fmt.Println(origin)
		for _, o := range origins {
			if o == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", "GET,POST")
				c.Header("Access-Control-Allow-Headers", "Content-Type")
			}
		}

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}

		c.Next()
	})

	r.GET("/user", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": "xiaoyesensen",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "This is an SSR Page",
		})
	})

	r.Run(":3000")
}
