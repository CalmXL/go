package main

import (
	"configuration/6_viper_3/config"
	"fmt"
)

/*
*

	  用户环境变量 - 软件运行时使用的
	  系统环境变量 - 系统运行时使用的

		%USER_NAME% => xiaoyesensen


	  MAC 环境变量
		.bash_profile 中

		终端
		export USER_NAME=xiaoyesensen
		export JAVA_HOME = /usr/java/jdk-1.8

		open ~/.bash_profile 打开文件
		export USER_NAME=xiaoyesensen
		export JAVA_HOME = /usr/java/jdk-1.8

		生效的命令
		source ~/.bash_profile

		开发环境下要使用环境变量
		Development     开发环境        DEBUG
		Production      生产环境


		环境与配置的问题
		DEV - port 3300
		DEBUG - port 8000
		PROD - port 8800
*/
func main() {
	cfg, _ := config.Initialize()

	fmt.Println(cfg.DBConfig.Port)
	fmt.Println(cfg.DBConfig.DBName)
}
