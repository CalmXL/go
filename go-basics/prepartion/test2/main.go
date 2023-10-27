package main

import "fmt"

/*
	* go build main.go 编译并生成可执行文件

	* go run main.go  运行 go 程序
		无论有没有可执行文件，都会先编译(二进制)再运行程序

	* go mod init test2
		初始化项目： 创建 go.mad -> 包管理工具

	* go get github.com/go-redis/redis 下载并安装包与其依赖项

	go install 编译与安装包

	go clean
		删除源码包和关联源码包里编译生成的文件

	go doc fmt.Println
		查看文档
	go env 查看 go 环境变量
	go fmt 格式化文档
	* go version 查看当前 Go 当前版本
*/
func main() {
	fmt.Println("Hello World")
}
