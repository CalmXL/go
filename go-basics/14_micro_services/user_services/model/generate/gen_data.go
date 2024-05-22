package main

import (
	"14_micro_services/user_services/model"
	"14_micro_services/user_services/util"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func generateUsers(db *gorm.DB) {

	suffix := ""
	for i := 0; i < 30; i++ {

		if i < 10 {
			suffix = "0" + strconv.Itoa(i)
		} else {
			suffix = strconv.Itoa(i)
		}

		fmt.Println(util.GeneratePassword("a12345678"), len(util.GeneratePassword("a12345678")))
		user := model.User{
			MobileNumber: "138000000" + suffix,
			Password:     util.GeneratePassword("a12345678"),
			NickName:     "user-" + suffix,
		}

		db.Save(&user)
	}
}
