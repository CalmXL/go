package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sale_system/controller"
	"sale_system/service"
)

func Sale(db *gorm.DB, r *gin.RouterGroup) {
	sale := controller.Sale{
		S: &service.Sale{
			DB: db,
		},
	}

	{
		r.Group("/category").
			GET("/list/:page_size/:page_number", sale.GetList).
			GET("/one/:id", sale.GetOne).
			POST("/add", sale.Add).
			POST("/update", sale.Update).
			POST("/delete", sale.Delete).
			POST("/searchKey", sale.SearchWithKeyword).
			POST("/searchDate", sale.SearchWithDate)

	}
}
