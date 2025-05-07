package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
  REST: Representational State Transfer 表现层状态转化

	ful => full
	Restful: 用 url 定位资源的 HTTP 标准风格
		GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS
	Resource: 资源 => 数据、文件、流
	Representation: 表现层 => 数据、文件或者流的类型 JSON, MP4, PNG
	State Transfer: 状态的变化 => GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS

	URI => 统一资源标识符，资源定位的标识
	URL => 同一资源定位符，寻找资源的具体路径， 远程地址，HTTP 网址， ws 长连接地址， FTP 资源地址

	Restful API => 使用 Restful 风格的 API, 也是遵循 HTTP 协议的一种接口设计风格

	CRUD | CURD => Create Read Update Delete
*/

type InfoBody struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint64 `json:"age"`
}

var infos = []InfoBody{
	{
		Id:   1,
		Name: "Mike",
		Age:  25,
	},
	{
		Id:   2,
		Name: "Nike",
		Age:  17,
	},
	{
		Id:   3,
		Name: "Mark",
		Age:  29,
	},
}

/**
	Access-Control-Allow-Origin
  同源策略
	浏览器的策略 => 浏览器对于资源请求的一种策略
	http://localhost:3000/info/2 源

	协议: http
	子域: www
	端口: 80

	这三个东西必须一致才是同源，才不会受到同源策略的影响
	同源策略引起的是跨域的问题
*/

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

func main() {
	r := gin.Default()
	r.Use(Cors())
	// GET 查询数据
	// http://localhost:8080?id=2
	r.GET("/info", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Query("id"), 10, 64)

		var _info InfoBody

		for _, info := range infos {
			if info.Id == id {
				_info = info
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": _info,
		})
	})

	r.GET("/info/:id", func(c *gin.Context) {

		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		var _info InfoBody

		for _, info := range infos {
			if info.Id == id {
				_info = info
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": _info,
		})
	})

	// 增加数据
	r.POST("/create", func(c *gin.Context) {
		var body InfoBody

		c.ShouldBindJSON(&body)

		newBody := InfoBody{
			Id:   infos[len(infos)-1].Id + 1,
			Name: body.Name,
			Age:  body.Age,
		}

		infos = append(infos, newBody)

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": newBody,
		})
	})

	// 修改数据
	r.PUT("/update", func(c *gin.Context) {

		name := c.PostForm("name")
		age, _ := strconv.ParseUint(c.PostForm("age"), 10, 64)
		id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

		for index, info := range infos {
			if info.Id == id {
				// info.Name = name
				// info.Age = age
				infos[index].Name = name
				infos[index].Age = age
				fmt.Println(info)
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": infos,
		})
	})

	// 更新数据一部分
	r.PATCH("update/age", func(c *gin.Context) {
		age, _ := strconv.ParseUint(c.PostForm("age"), 10, 64)
		id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

		for index, info := range infos {
			if info.Id == id {
				infos[index].Age = age
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": infos,
		})
	})

	// 删除数据
	r.DELETE("/delete", func(c *gin.Context) {
		var body InfoBody

		c.ShouldBindJSON(&body)

		for index, info := range infos {
			if info.Id == body.Id {
				infos = append(infos[:index], infos[index+1:]...)
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": infos,
		})
	})

	r.HEAD("/head", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, nil)
	})

	r.OPTIONS("/methods", func(c *gin.Context) {
		headers := "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS"
		c.JSON(http.StatusOK, headers)
	})

	r.Run(":3000")
}
