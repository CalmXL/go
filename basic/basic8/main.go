package main

import (
	"fmt"
	"strconv"
)

/*
	与字符串之间的转换
*/

func main() {
	// strconv string conversion
	/*
		 	con com co => come 一起，来
			vers => 相反、转换、变动
	*/

	// string类型 => int类型
	//var str = "123"
	//var str = "abc"
	// a => string
	// to => 转换
	// i => int
	// 字符串转其他类型, 有可能不成功
	//a, err := strconv.Atoi(str)
	//
	//if err != nil {
	//	fmt.Println(a)
	//}
	//
	//fmt.Println(a)

	// ***************************************
	// int类型 => 字符串类型
	//var a = 123
	//str := strconv.Itoa(a)
	//fmt.Println(str)

	// ***************************************
	// string 类型 => float 类型
	// 转换用 32 位， 返回的是 float64 的类型
	//f, err := strconv.ParseFloat("3.1415", 32)

	// string 类型 => int 类型
	// 把 "100" 看成 2进制进行 10 进制解析
	// 如果需要进制转换，就用 parseInt， 如果不需要，就直接使用 Atoi
	//i, errI := strconv.ParseInt("100", 2, 64)

	// string 类型 => 布尔类型
	//b, errB := strconv.ParseBool("true")
	//b, errB := strconv.ParseBool("1")

	// 除了 0 和 1 都会报错
	// strconv.ParseBool: parsing "123": invalid syntax
	//b, errB := strconv.ParseBool("123")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//if errI != nil {
	//	fmt.Println(errI)
	//}
	//
	//if errB != nil {
	//	fmt.Println(errB)
	//}
	//fmt.Println(f) // => 3.1414999961853027
	//fmt.Println(i) // => 4
	//fmt.Println(b) // => true

	// *******************************************
	// bool类型 => string类型
	a := strconv.FormatBool(true)

	// float类型 => string类型
	//b := strconv.FormatFloat(3.1415926, 'f', -1, 64)
	b := strconv.FormatFloat(3.1415926, 'f', 4, 64)

	// int 类型 => string 类型
	// 100 作为十进制 进行二进制转换输出
	c := strconv.FormatInt(100, 2) // => 1100100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
