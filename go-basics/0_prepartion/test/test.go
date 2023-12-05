package main

import "fmt"

/**
  package: 包
  package main
  同一个文件夹下，就不需要在引用
*/

/**
	func -> function 关键字
	声明一个函数
    函数名 main (不可变)
	() 参数列表
	{} 函数体
*/

/**
  	Go 不需要打分号的
	Go官方不推荐在一个文件夹下有多个 main
	一个文件下不推荐有多个文件声明 main
	一个文件内不允许有多个 main 声明
*/
func main() {
	//println("hello World!") 不推荐使用
	// Println => print line 输出一行
	fmt.Println("Hello world")
}
