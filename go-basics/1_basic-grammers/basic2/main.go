package main

import "fmt"

/*
	变量是干什么的？
		1. 变量最早是数学工具
			 当我们在计算中不确定值是什么的时候
       用变量进行暂时的替代（抽象）以完成运算
			 且变量是可变的不确定的量
       不同的运算中，变量所表示的最终值不同
       所以变量是可以变的量

		2. 变量在计算机中的概念
			1. 表示值的抽象工具
					x = 1 => 1 赋值给 x (= 赋值符号)
			2. 存储计算结果的工具
					y = x + 1 => 2

			在内存中变现的方式

			标识符			内存地址			值
				x				0x0001			1
				y				0x0002	    2
				z				0x0003      3

			0x0001 = 1
			0x0002 = 2
			0x0003 = 0x0001 + 0x0002

			x = 1
			y = 2
			z = 3

			总结：变量就是方便计算机语言编程时
				   从内存中取值和向内存中存值的工具
*/

/*
	变量的声明格式
		关键字			变量标识符	变量类型	赋值		值
		var 			x					int     =     1
		variable

		1. 变量声明可以不初始化
		2. 变量初始化可以在变量声明之后进行
		3. 变量赋值或初始化不是使用变量
		4. 值必须与变量声明时给出的类型保持一致

	1. 变量必须在声明之后使用
    2. 变量声明后必须使用
	3. 同一作用域下，不能重复声明
	4. 多个变量可以使用单一声明法
	5. 多个变量可以批量声明
	6. 类型推断
	7. 可以使用短变量语法（只能使用在函数内部，不能使用在函数外部）
	8. 短变量可以批量声明赋值
	9. 可以使用匿名变量
*/

/*
	Go 语言的函数外部是不能使用赋值语句的
	只能出现声明语句
*/

// x := 3
var x = 3 // 声明语句
func main() {
	//1. 变量声明可以不初始化
	//var x int
	// 变量初始化
	//x = 1
	//
	//fmt.Println(x)

	// 先声明后使用
	//fmt.Println(x)
	//var x int

	// 变量必须使用
	//var x int

	// 正常声明多个变量
	//var x int = 1
	//var y string = "123"
	//var z bool = true
	//
	//fmt.Println(x, y, z)

	// 单一声明
	//var x, y, z int
	//x = 1
	//y = 2
	//z = 3
	//
	//fmt.Println(x, y, z)

	//var x, y, z = 1, "abc", true
	//fmt.Println(x, y, z)

	// 批量声明
	//var (
	//	x int    = 1
	//	y string = "abc"
	//	z bool   = true
	//)

	//var (
	//	x int
	//	y string
	//	z bool
	//)
	//
	//x = 1
	//y = "123"
	//z = true

	// 类型推断
	/**
	int x = 1
	String y = abc
	boolean = true
	*/

	//var x = 1
	//var y = "abc"
	//var z = true

	// 短变量 :=
	/*
		var x int 声明语句
		x = 1 赋值语句
	*/
	//x := 1
	//y := "abc"
	//z := true

	// 短变量 批量声明赋值
	//x, y, z := 1, "abc", true

	// 匿名变量
	// _ 匿名变量，可以不使用
	x, _, z := getXYZ()
	//fmt.Println(x, y, z)
	fmt.Println(x, z)
}

func getXYZ() (int, string, bool) {
	return 1, "abc", true
}
