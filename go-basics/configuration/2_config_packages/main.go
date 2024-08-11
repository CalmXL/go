package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// if err := godotenv.Load("./config/app1.env", "./config/app2.env"); err != nil {
	// 	log.Fatalf("读取环境变量出错: %s", err.Error())
	// }

	// if err := godotenv.Load("./config/app1.yaml", "./config/app2.yaml"); err != nil {
	// 	log.Fatalf("读取环境变量出错: %s", err.Error())
	// }

	// if err := godotenv.Load("./config/app1.json", "./config/app2.json"); err != nil {
	// 	log.Fatalf("读取环境变量出错: %s", err.Error())
	// }

	// envMap, err := godotenv.Read("./config/app1.env", "./config/app2.env")
	//
	// if err != nil {
	// 	log.Fatalf("读取环境变量失败: %s", err.Error())
	// }
	//
	// fmt.Println(envMap["APP_ENV"])

	// appEnv := os.Getenv("APP_ENV")
	// baseUrl := os.Getenv("BASE_URL")
	//
	// fmt.Println(appEnv, baseUrl)

	// ****************************************************
	/**
	  viper 支持几乎所有的配置型文件
	*/
	viper.SetConfigFile("./config/app.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}

	appEnv := viper.Get("APP_ENV")
	baseUrl := viper.Get("BASE_URL")

	fmt.Println(appEnv, baseUrl)
}
