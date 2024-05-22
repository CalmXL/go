package calculator

/*
  1. func -> function 函数的声明
	2. Plus -> 函数名称（需要公共调用的方法需要大坨峰）
						 GetName
  3. 类型定义 -> Go 语言定义放在变量后面
  4. a 与 b 都是函数参数，在调用（Plus(1, 2)）, 调用传入实际参数
	5. 返回值类型定义 -> 参数括号后面
	6. GO 中不需要打分号
*/

func Plus(a int, b int) int {
	PrintOver()
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}
