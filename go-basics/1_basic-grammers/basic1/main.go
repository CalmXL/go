package main

import (
	"basic1/calculator"
	"fmt"
)

/*
	Identifiers 标识符
 	Go 的程序员自己可以决定的内筒全部都叫做标识符
	变量、常量、函数、结构体、接口

	标识符命名方式：
		camelCase: 驼峰命名法 - helloWorld | getCurrentTime
		PascalCase: 帕斯卡命名法- HelloWorld | GetCurrentTime
		snake_case: 蛇形命名法 - hello_world | HELLO_WORLD
		kebab-case: 中横线命名法 - hello-world
		space case: 空格命名法 编程语言使用 x

	Go 语言中使用标识符的元素：（除开下划线以外的所有符号都不可以使用）
		1. 大写字母
		2. 小写字母
		3. 数字
		4. 下划线

	规则:
		1. 不能以用数字开头
		2. 不能带有_以外的符号
		3. 不允许有空格
		4. 严格区分大小写
		5. 语义清晰明了、简介 getCurrentTime
		6. 关键字 与 预定义标识符不能作为标识符 import package

	GO 中的命名规范
		项目名:
			1. 中横线隔开
			2. 全部小写字母

		包名:
			1. 包名与目录名保持一致
			2. 全部小写
			3. 不能是用下划线
			4. 不能用标准库的名称 os/fmt

		模块名：
			1. 全部小写
			2. 单词用下划线隔开 User_info.go

		常量： HTTP_PORT
			1. 全部大写
			2. 用下划线隔开

		结构体：遵循变量的命名规则

		接口名: er 后缀，遵循变量的命名规则

		函数名和普通的变量: 严格区分大小写
			1. 任何需要在包外暴露被使用的标识符都必须以大写字母开头
					fmt.Println() -> public
			2. 任何不需要在包外暴露使用的标识符都必须以小写字母开头
					user.getName() -> private

*/

func main() {
	//var 123a int = 123 x
	//var $134 int = 123 x
	//var myTitle string = "my title1"
	//var MyTitle string = "My Title2"
	//
	//fmt.Println(myTitle, MyTitle)
	//var import int = 13 x

	result := plus(1, 2)
	result2 := calculator.Subtract(7, 2)
	result3 := multiply(3, 4)
	myName := calculator.MY_NAME

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(myName)

}

func plus(a int, b int) int {
	return a + b
}
