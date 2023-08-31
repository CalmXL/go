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

	a := "abc"
	b := "bcd"
	res := a + b
	fmt.Println(res) // => abcbcd
}
