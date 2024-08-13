package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
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
	// viper.SetConfigFile("./config/config.yaml")
	// viper.SetConfigFile("./config/config1.json")
	//
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("读取配置文件失败: %s", err)
	// 	return
	// }
	//
	// appEnv := viper.Get("app_env")
	// fmt.Println(appEnv)

	/**
	  通过 AddConfigPath SetConfigName SetConfigType ReadInConfig 读取配置文件内容

		先以文件名优先级进行匹配
	*/
	// viper.AddConfigPath("./config")
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	//
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// 	return
	// }
	// appEnv := viper.Get("app_env")
	// fmt.Println(appEnv)

	/**
	4. 别名
	*/
	// viper.Set("app_env", "production")
	// viper.RegisterAlias("env", "app_env")
	// env := viper.Get("env")
	//
	// fmt.Println(env)

	/**
	  5. 设置默认值
	*/

	// viper.SetDefault("app_env", "development")
	// viper.SetDefault("db", map[string]any{
	// 	"host":     "127.0.0.1",
	// 	"port":     3306,
	// 	"user":     "root",
	// 	"password": "xL1210...",
	// 	"dbname":   "study",
	// })
	//
	// fmt.Println(viper.Get("app_env"))
	// fmt.Println(viper.Get("db"))

	/**
	  6.写入文件 和 监听配置变化
	*/
	viper.SetConfigFile("./config/db.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e.Name)
	})

	viper.Set("db", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     3306,
		"user":     "root",
		"password": "xL1210...",
		"dbname":   "study",
	})

	viper.Set("db.type", "mysql")
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error writing config, %s", err)
		return
	}

	if err := viper.SafeWriteConfigAs("./config/database.yaml"); err != nil {
		log.Fatalf("Error writing config, %s", err)
	}
}
