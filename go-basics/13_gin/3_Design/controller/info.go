package controller

import (
	"13_gin/3_Design/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Info struct {
}

func (info *Info) GetInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Query("id"), 10, 64)

	var _info data.InfoBody

	for _, info := range data.Infos {
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
}

func (info *Info) GetInfo2(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var _info data.InfoBody

	for _, info := range data.Infos {
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
}

func (info *Info) Create(c *gin.Context) {
	var body data.InfoBody

	c.ShouldBindJSON(&body)

	newBody := data.InfoBody{
		Id:   data.Infos[len(data.Infos)-1].Id + 1,
		Name: body.Name,
		Age:  body.Age,
	}

	data.Infos = append(data.Infos, newBody)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": newBody,
	})
}

func (info *Info) Update(c *gin.Context) {

	name := c.PostForm("name")
	age, _ := strconv.ParseUint(c.PostForm("age"), 10, 64)
	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

	for index, info := range data.Infos {
		if info.Id == id {
			// info.Name = name
			// info.Age = age
			data.Infos[index].Name = name
			data.Infos[index].Age = age
			fmt.Println(info)
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data.Infos,
	})
}

func (info *Info) UpdateAge(c *gin.Context) {
	age, _ := strconv.ParseUint(c.PostForm("age"), 10, 64)
	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

	for index, info := range data.Infos {
		if info.Id == id {
			data.Infos[index].Age = age
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data.Infos,
	})
}

func (info *Info) Delete(c *gin.Context) {
	var body data.InfoBody

	c.ShouldBindJSON(&body)

	for index, info := range data.Infos {
		if info.Id == body.Id {
			data.Infos = append(data.Infos[:index], data.Infos[index+1:]...)
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data.Infos,
	})
}

func (info *Info) Head(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, nil)
}

func (info *Info) GetMethods(c *gin.Context) {
	headers := "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS"
	c.JSON(http.StatusOK, headers)
}
