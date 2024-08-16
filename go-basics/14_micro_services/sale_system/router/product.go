package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sale_system/controller"
	"sale_system/service"
)

func Product(db *gorm.DB, r *gin.RouterGroup) {
	product := controller.Product{
		S: &service.Product{
			DB: db,
		},
	}

	{
		r.Group("/category").
			GET("/list/:page_size/:page_number", product.GetList).
			GET("/one/:id", product.GetOne).
			POST("/add", product.Add).
			POST("/update", product.Update).
			POST("/delete", product.Delete).
			POST("/search", product.Search)

	}
}
