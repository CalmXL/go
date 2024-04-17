package main

import "fmt"

// type MyInt int
// type MyInt2 = int

func main() {
	// var a MyInt = 1
	// var b int = 2
	// var c MyInt2 = 3
	//
	// // 无效运算: a == b(类型 MyInt 和 int 不匹配)
	// // fmt.Println(a == b)
	//
	// res := int(a) + b
	// fmt.Println(res)
	//
	// res2 := b + c
	// fmt.Println(res2)
	
	// type Account float64
	//
	// var accountBalance Account
	// var income Account = 123.12
	// total := accountBalance + income
	//
	// fmt.Println(total) // => 123.12
	
	// interface{}
	// type StudentMap map[string]interface{}
	//
	// var studentMap StudentMap = map[string]interface{}{
	// 	"name": "张三",
	// 	"age":  18,
	// }
	//
	// fmt.Println(studentMap)
	
	var slice ComputeSlice = []int{1, 2}
	
	fmt.Println(slice.compute("PLUS"))
	fmt.Println(slice.compute("MINUS"))
	fmt.Println(slice.compute("MULTIPLY"))
	fmt.Println(slice.compute("DIVISION"))
	
}

type ComputeSlice []int

// 给自定义类型提供方法
// 给自定义类型的所欲数据提供方法
func (cs ComputeSlice) compute(method string) int {
	switch method {
	case "PLUS":
		return Plus(cs[0], cs[1])
	case "MINUS":
		return Minus(cs[0], cs[1])
	case "MULTIPLY":
		return Multiply(cs[0], cs[1])
	case "DIVISION":
		return Division(cs[0], cs[1])
	default:
		return 0
	}
}

func Plus(a int, b int) int {
	return a + b
}
func Minus(a int, b int) int {
	return a - b
}
func Multiply(a int, b int) int {
	return a * b
}
func Division(a int, b int) int {
	return a / b
}
