package main

import (
	"fmt"
)

func main() {
	// var a int8 = 123
	// 无法将 'a' (类型 int8) 用作类型 uint8
	// var b uint8 = a

	//	var a int8 = 123
	//	var b uint8 = uint8(a)
	//	fmt.Println(b)

	// +++++++++++++++++++++++++++++++++++++++++++++

	// var a int8 = -123
	// 无法将 'a' (类型 int8) 用作类型 uint8
	// var b uint8 = a

	// var a int8 = -123
	// var b uint8 = uint8(a)
	// // 256 - 123 = 133
	// fmt.Println(b) // => 133

	// +++++++++++++++++++++++++++++++++++++++++++++

	// var a int32 = 123
	// //无法将 'a' (类型 int32) 用作类型 int64
	// var b int64 = a
	//
	// var a int32 = 123
	// var b = int64(a)
	// fmt.Println(b)

	// +++++++++++++++++++++++++++++++++++++++++++++
	// var a int64 = 123
	// var b = int32(a)
	// fmt.Println(b)
	//
	// +++++++++++++++++++++++++++++++++++++++++++++
	//
	// var a = 3.1415
	// var b = float32(a)
	// fmt.Println(b)
	//
	// +++++++++++++++++++++++++++++++++++++++++++++
	//
	// var a = 123
	// var b = float64(a)
	// fmt.Println(b) // => 123
	//
	// +++++++++++++++++++++++++++++++++++++++++++++
	// 自定义的类型 不能用作默认类型
	// type typeAmount float64
	// var a = 1000.12423
	// // 无法将 'a' (类型 float64) 用作类型 typeAmount
	// var b = typeAmount(a)
	// fmt.Println(b)
	//
	// ++++++++++++++++++++++++++++++++++++++++++++
	//
	type typeAmount float64
	var amount typeAmount = 1000.2334
	var newAmount = 100.23
	//
	// 无效运算: amount - newAmount(类型 typeAmount 和 float64 不匹配)
	// num := amount - newAmount
	// num := float64(amount) - newAmount
	// fmt.Println(num) // => 900.0033999999999

	num2 := amount - typeAmount(newAmount)
	fmt.Println(num2) // => 900.0033999999999

	// ++++++++++++++++++++++++++++++++++++++++++++

	// 获取类型
	/*
		reflect.TypeOf()
	*/
	//a := true
	//b := 123
	//c := "abc"
	//d := 3.14
	//
	//aType := reflect.TypeOf(a)
	//bType := reflect.TypeOf(b)
	//cType := reflect.TypeOf(c)
	//dType := reflect.TypeOf(d)
	//
	// aaType := reflect.ValueOf(a).Kind()
	// bbType := reflect.ValueOf(b).Kind()
	// ccType := reflect.ValueOf(c).Kind()
	// ddType := reflect.ValueOf(d).Kind()
	//
	// fmt.Println(aType, bType, cType, dType)     // => bool int string float64
	// fmt.Println(aaType, bbType, ccType, ddType) // => bool int string float64
}
