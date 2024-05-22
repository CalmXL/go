package main

import (
	"12_gorm/db"
	"12_gorm/model"
	"fmt"
)

func main() {
	db, err := db.Connect("golang")

	if err != nil {
		fmt.Println(err.Error())
	}

	// var s model.Student
	// 获取数据表的第一条 id 正序
	// db.First(&s)
	// 获取数据表的最后一条 id 倒序
	// db.Last(&s)

	// db.Take(&s)

	// s := map[string]interface{}{}
	// // 没有指定具体的数据模型的情况下，必须使用 model 方法
	// db.Model(&model.Student{}).First(&s)
	//
	// fmt.Println(s)

	var s model.Student
	db.First(&s, 16)
	fmt.Println(s)
}
