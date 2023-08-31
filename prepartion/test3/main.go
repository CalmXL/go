package main

import (
	"fmt"
	"test3/calculator"
)

/*
	项目名称: test3
	包名称: calculator
	模块名称: logic
*/

/*
	1. 创建项目： 创建文件夹 | Goland 新建项目
	2. 初始化项目: go mod init test3
	3. 创建包: 新建 calculator 文件夹
	4. 创建模块: 新建 logic.go 文件
*/

/*
	1. 有main 方法的文件， 必须有 package main 声明
  2. 一个文件中只允许有一个 main
	3. 一个文件下，同级目录不推荐有多个 main 方法
	4. 当运行 go 文件时, main 方法会被自动调用
	5. 一个包内部调用，不需要包名引导
			logic.go -> words.go -> PrintOver 不需要 calculator
	6. 不同包之间调用方法，需要包名引导
*/

// main 方法，GO 程序的入口方法
func main() {
	//var result int = calculator.Plus(1, 2)
	//var result2 int = calculator.Subtract(5, 1)
	//var result3 int = calculator.Multiply(3, 4)
	//var result4 int = calculator.Divide(10, 2)

	result := calculator.Plus(1, 2)
	result2 := calculator.Subtract(5, 1)
	result3 := calculator.Multiply(3, 4)
	result4 := calculator.Divide(10, 2)

	fmt.Println("result: ", result)
	fmt.Println("result2: ", result2)
	fmt.Println("result3: ", result3)
	fmt.Println("result4: ", result4)
	calculator.PrintOver()
}

//func main() {
//
//}
