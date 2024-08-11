package main

import (
	"github.com/spf13/viper"
	"log"
)

/**
  viper 是 GO 语言中使用最广泛的配置工具
	1. 环境的参数设置
	2. 项目的配置项
	3. 命令参数
*/

func main() {
	/**
	  1. 使用 Set 和 Get 方法进行配置的设置与读取
	*/
	// viper.Set("app_env", "development")
	// appEnv := viper.Get("app_env")
	//
	// fmt.Println(appEnv)

	/**
	  2. 使用 SetConfigFile 和 ReadInConfig 读取配置文件的内容
	*/
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %s", err)
		return
	}

	appEnv := viper.Get("app_env")
}
