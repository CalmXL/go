package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*

	    Must bind - 400
			Bind， BindJSON, BindXML, BindQuery, BindYAML
		Should bind
			ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML, ShouldBindHeader
*/
func main() {
	r := gin.Default()

	/**
	query 设置了没有作用， 必须使用 Name, Age, IsMarried 必须大写
	*/
	type Info struct {
		Name      string `json:"name" xml:"name" binding:"required"`
		Age       int32  `json:"age" xml:"age" binding:"required"`
		IsMarried bool   `json:"isMarried" xml:"isMarried" binding:"required"`
	}

	r.POST("/info", func(c *gin.Context) {
		var info Info

		// c.MustBindWith(&info, binding.JSON)

		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"msg": err.Error(),
		// 	})
		// 	return
		// }
		// if err := c.ShouldBindWith(&info, binding.JSON); err != nil {

		// if err := c.Bind(&info); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"msg": err.Error(),
		// 	})
		// 	return
		// }

		// c.ShouldBindXML(&info)

		err := c.ShouldBindQuery(&info)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, info)
	})

	err := r.Run()
	if err != nil {
		return
	}
}
