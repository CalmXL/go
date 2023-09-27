package main

import "fmt"

/*
	算数运算：
	+ 加法算数运算、字符串拼接、正号
	-	减法算数运算、负号
	* 乘法算数运算
	/ 除法算数运算
	% 取余算数运算
	++ 自加算数运算
	-- 自减算数运算
*/

func main() {
	//a := 3.14
	//b := 100
	//c := true
	//var d byte = 'a'
	//e := "bcd"
	//
	//fmt.Println(a, b, c, d, e)

	// 加法运算
	// 算数运算中： go 不存在隐式的类型转换
	// 效运算: a + b(类型 float64 和 int 不匹配)
	// 无效运算: a + b + c(类型 float64 和 bool 不匹配)
	//res := a + b + c + d + e

	// res := a + b x
	// res := a + float64(b) // => 103.14
	//res := int(a) + b // => 103
	//fmt.Println(res)
	//
	//fmt.Println(100 + 3.14) // => 103.14

	// *************************************

	//a := true
	//b := false

	// Invalid operation: a + b (the operator + is not defined on bool)
	// 无效运算: a + b (在 bool 中未定义运算符 +)
	//res := a + b

	// *************************************
	//a := true
	//b := 1

	// Cannot convert an expression of the type 'bool' to the type 'int'
	//res := int(a) + b

	// *************************************
	// 字符串拼接
	//a := "abc"cc
	//b := "bcd"
	//res := a + b
	//fmt.Println(res) // => abcbcd

	// *************************************
	//a := 1
	//b := -a
	//fmt.Println(b) // => -1

	// *************************************
	//a := "abc"
	//var b byte = 'd'
	//res := a + b x 类型不匹配

	// *************************************
	//var a byte = 'a' // -> uint8 97
	//var b byte = 'b' // -> uint8 98
	//
	//res := a + b
	//fmt.Println(res) // => 195

	// *************************************
	//a := 1
	//b := 2
	//
	//res1 := a + b
	//res2 := a - b
	//res3 := a * b
	//res4 := a / b // res4 int => 小数部分丢失
	//res5 := a % b // 1 / 2 => 0 , 1
	//
	//fmt.Println(res1, res2, res3, res4, res5)

	// *************************************
	//a := "abc"
	//b := "bcd"

	//res := a - b
	//res2 := a * b
	//res3 := a / b
	//
	//fmt.Println(res, res2, res3)

	// *************************************
	// uint8 => 0~255
	//var a byte = 'a'
	//var b byte = 'b'
	//
	//res := a - b     // 97 - 98 = -1 => 255
	//fmt.Println(res) // => 255

	// *************************************
	//var a = 0
	//var b = 0
	//var c = 1
	// panic: runtime error: integer divide by zero
	//fmt.Println(a / b)
	//fmt.Println(c / a)

	// *************************************
	// +Inf
	a := 1.0
	b := 0.0

	fmt.Println(a / b) // => +Inf
}
