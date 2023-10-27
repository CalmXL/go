package main

import "fmt"

/*
	1. 常量的概念：
		 1. 数学中的常量 （常数）
				恒定不变的自然数或定义数
				π 圆周率
				γ 欧拉常数
				i 虚数
				∞	无穷大

		 常量的作用：是用符号代替比较复杂的数值来表示参与运算的过程
		 数学中的常量值表示的是实用性的值，而不是操作性的值

		2. 计算机语言中的常量
			1. 沿用了数学中常数的基本特性
			2. 常量是被开发者定义的量
			3. 常量在初始化之后，是不可以被修改的
			4. 常量不可以在同一作用域下重复声明
		常量作用: 代替指的重复使用，且限制其值的改变
		原理：常量在编译时确定其值，运行时不可变其值

  2. 定义常量的方法
		关键字				标识符  	常量类型			赋值符号  	值
    const        PI			float32     =					3.14
		constant

	3. 注意事项
		1. 常量在初始化时必须声明初始化值
		2. 常量没有严格要求必须使用，但会警告， 不影响编译
		3. 常量值不可修改
		4. 常量也可以使用类型推断
		5. 常量可以批量声明
		6. 常量可以单一声明
*/

func main() {
	// 声明常量
	// 未使用的常量 'PI' (只是警告，编译不受影响)
	//const PI float32 = 3.14

	// pi declared but not used (编译不通过)
	//var pi float32 = 3.14

	//PI = 3.15

	// 类型推断
	//const PI = 3.14

	// 批量声明
	//const (
	//	name    = "xiaoyesensen"
	//	age     = 37
	//	married = true
	//)
	//
	//fmt.Println(name, age, married)

	// 单一声明
	//const name, age, married = "xiaoyesensen", 37, true
	//fmt.Println(name, age, married)

	/*
		iota 不是英文单词的缩写
		希腊字母中的 I 的读音 => [yota]
		英文单词读音 => [ai'outa]
		英文单词的意思 => 微小的量

		前言：
			批量声明常量的时候，第一个常量必须要初始化。后续的常量可以不初始化
			h后续常量在编译时自动赋值第一个常量的值

			总结： 后续没有初始化值的常量，在编译时，会赋值为最后一个常量的值
	*/

	//const a, b, c, d, e = 0, 1, 2, 3, 4

	//const (
	//	a = 0
	//	b
	//	c
	//	d
	//	e
	//)

	/*
			iota 在 GO 程序编译时可以修改常量值，默认起始值为 0
			常量每声明一个， iota 迭代一次 (+1)
		  运用 iota 相当于使用常量迭代器来初始化常量
	*/
	//const (
	//	a = iota
	//	b
	//	c
	//	d
	//	e
	//)

	const (
		a = 0
		b
		c = iota
		d
		e
	)
	fmt.Println(a, b, c, d, e)

	//const (
	//	a = iota
	//	_
	//	c
	//	_
	//	e
	//)
	//fmt.Println(a, c, e) // => 0 2 4

	type Weekday int

	const (
		SUN Weekday = iota
		MON
		TUE
		WED
		THU
		FRI
		SAT
	)

	fmt.Println(SUN, MON, TUE, WED, THU, FRI, SAT) // => 0 1 2 3 4 5 6

}
