package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  中间件:
		函数中间的某个函数

	func test1 (c *gin.Context) {
		// todo1
		c.Next() // next 是对下一个中间件函数的引用
		// todo2
	}
	func test2 (c *gin.Context) {
		// todo3
		c.Next()
		// todo4
	}
	func test3 (c *gin.Context) {
		// todo5
		c.Next()
		// todo6
	}

	// 洋葱模型
	test1() => todo1 => next()
	test2() => todo3 => next()
	test3() => todo5 => next()
	todo6 => todo4 => todo2

*/

func test1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo1
		fmt.Println("todo1")
		c.Next() // next 是对下一个中间件函数的引用
		// todo2
		fmt.Println("todo2")
	}
}

func test2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo3
		fmt.Println("todo3")
		// c.Next() // 暂停当前中间件函数的执行，执行下一个中间件函数
		c.Abort() // 不走下一个中间件
		// return  // 不再执行当前中间件函数接下来的程序
		// todo4
		fmt.Println("todo4")
	}
}

func test3() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo5
		fmt.Println("todo5")
		c.Next()
		// todo6
		fmt.Println("todo6")
	}
}

func loginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if len(username) < 6 {
			c.JSON(http.StatusOK, gin.H{
				"code": "1001",
				"msg":  "Invalid username length",
				"data": nil,
			})
			c.Abort()
			return
		}

		if len(password) < 8 {
			c.JSON(http.StatusOK, gin.H{
				"code": "1001",
				"msg":  "Invalid password length",
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

var origins = []string{
	"http://localhost:5173",
	"http://localhost:3000",
}

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

func main() {
	r := gin.Default()

	// 如果你用了use, 那么访问每个接口之前都要调用只写中间件函数
	// 一般用在所有接口都要执行的函数逻辑上
	// r.Use(test1(), test2(), test3())

	r.Use(Cors())

	// 单独的中间件函数
	// r.GET("/test", test1(), test2(), test3(), func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, nil)
	// })

	r.GET("/test1", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.POST("/login", loginCheck(), func(c *gin.Context) {
		userName := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": gin.H{
				"username": userName,
				"password": password,
			},
		})
	})

	r.Run()
}
