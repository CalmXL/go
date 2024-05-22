package main

import "fmt"

/**
  这就是一个函数
	也是一个 main 函数
  也是 Go 程序的入口函数 (自执行一次)
*/

/**
  func 函数声明关键字
	main：函数名
	(): 参数容器 (形式参数)
	{}: 函数体 函数逻辑容器
*/

/**
  函数(function): Go 语言中的一等公民

	1. 函数存在的目标
		a. 对特定程序的进行封装
    b. 对可复用的程序封装
		c. 函数是一等公民的特征，编写强大的函数接口

	2. 函数一等公民的特征
		a. 函数可以进行参数传递 - (回调特性)
		b. 函数可以作为返回值抛出 - (闭包特性)
		c. 函数可以作为值进行变量赋值 - (值特性)
		d. 函数可以实现接口 (满足接口特性)

	3. main 函数的特征
		a. main 函数就是一个普通的 Go 函数，可以在任何地方调用执行
		b. main 函数在 Go 程序中自动执行一次
		c. main 函数无参数 (main 函数不得有实参和返回值)
		d. main 函数不能有返回值

	4. 函数与作用域的关系

		什么是作用域？
			1. 作用域就是变量可以访问的范围
			2. 作用域就是变量的访问容器

		a. main 函数外有一个在 一个 Go 文件内，所有函数都能访问的作用域 - 全局作用域，
			 必须使用 var 声明不能使用 :=
		b. 每个函数内是一个局部(函数)作用域
				在自己作用域中访问变量，自己有的变量，不会访问外部同变量名的变量
				函数作用域 -- 可访问 --> 外部作用域
		c. 函数内部也可以声明函数，函数作用域是可以嵌套的
				函数内部只能定义函数表达式，不能声明
		d. 内部作用域可以访问外部作用域，外部作用域无法访问内部作用域的

		作用域链：
			全局作用域 -> main 函数作用域 -> test 函数作用域

   5. 函数的种类
		  a. 函数的声明
				只能出现在全局作用域，函数内部不能进行具名函数声明
				func test () {}
      b. 函数的表达式
				匿名函数声明赋值给一个变量的表达式
				test := func () {}

*/

var a = 1
var b = 100

func test4() {
	// main()
}

var test6 = func() {}

func main() {
	// a, b := 1, 2
	// c, d := 3, 4
	// // res := a + b
	// // res2 := c + d
	//
	// // a, b 是实际参数 (实参)
	// res1 := plus(a, b)
	// res2 := plus(c, d)
	//
	// fmt.Println(res1, res2)

	// closure := test2()
	// closure()

	// funcValue := test3
	//
	// funcValue()

	// fmt.Println(a)

	test := func() {
		a := 3
		fmt.Println("test: ", a)
		fmt.Println("test b: ", b)
	}

	test()

	fmt.Println("main: ", a)
}

// 函数名更具备语义化
// func plus(num1 int, num2 int) int {
func plus(num1, num2 int) int {
	res := num1 + num2
	// 返回值
	return res
}

// cb => callback 回调函数 => 函数参数
func test1(cb func()) {
	cb()
}

func test2() func() {
	/**
	  函数内部定义的的函数就是闭包
	*/
	return func() {
		fmt.Println("Closure")
	}
}

func test3() {
	fmt.Println("test3")
}
