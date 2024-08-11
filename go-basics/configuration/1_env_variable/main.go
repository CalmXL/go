package main

import (
	"fmt"
	"log"
	"os"
)

/**
环境变量：Environment variable
	软件运行环境中所需的一些参数
	软件在运行过程中，是一定有些必要的配置项，系统或者软件一般是用环境变量来实现的
	1. 软件运行时需要的路径指示
	2. 软件运行时的环境类型
	3. 软件运行中必要的配置

	os 包 => Go 语言当中针对于操作系统的包
		Setenv Getenv Lookupenv  unsetenc


	配置文件
		.env  专门给环境变量配置的文件类型
		.json 专门给配置做集中管理的格式化文件
		.yaml 一个书写格式比较简单的配置文件格式
*/

func main() {
	if err := os.Setenv("APP_ENV", "development"); err != nil {
		log.Fatalf("设置环境变量失败：%s", err)
		return
	}

	// 通过 get 方式无法直接得到一个是否成功获取的结果
	// abc := os.Getenv("APP_ABC")
	// fmt.Println("abc: ", abc)

	// 不近能够返回环境变量值，也能直接得知获取成功或失败
	// abc, ok := os.LookupEnv("abc")
	//
	// if !ok {
	// 	log.Fatal("获取 abc 失败")
	// 	return
	// }
	//
	// fmt.Println(abc)

	appEnv, ok := os.LookupEnv("APP_ENV")

	if !ok {
		log.Fatal("获取应用环境变量失败")
		return
	}

	fmt.Println(appEnv)

	if err := os.Unsetenv("APP_ENV"); err != nil {
		log.Fatalf("删除应用环境变量失败: %s", err.Error())
	}

	appEnv2 := os.Getenv("APP_ENV")
	fmt.Println(appEnv2)
}
