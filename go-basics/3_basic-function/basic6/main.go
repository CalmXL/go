package main

import "fmt"

/**
  自定义类型
	类型别名
*/

// 自定义类型
// 自定义类型是一种新的类型，与内置类型是不同类型
// 你使用 int 定义了一个新的类型叫做 MyInt
// type MyInt int

// MyInt 类型别名: 起到语义化的作用
// MyInt 是一个 int  类型的别名，本质上两个标识是一个类型(int)
// type MyInt = int

func main() {

	// var a MyInt = 1
	// var b int = 1
	//
	// var a rune = 1
	// var b int32 = 1
	//
	// // 无效运算: a == b(类型 myInt 和 int 不匹配)
	// fmt.Println(a == b)

	// ---------------------------------------------

	// 可以创建一个运算隔离墙
	// type Account float64
	//
	// var accountBalance Account
	// var income Account = 123.12
	// total := accountBalance + income

	// ---------------------------------------------

	// type StudentMap map[string]interface{}
	//
	// var studentMap StudentMap = map[string]interface{}{
	// 	"name": "张三",
	// 	"age":  18,
	// }

	// ---------------------------------------------
	// res := plus(1, 2)
	// fmt.Println(res)

	var slice ComputeSlice = []int{1, 2}

	/**
	  通过 (cs ComputeSlice) 给  ComputeSlice 类型的所有数据都提供了加减乘除的方法
	   所有该类型的方法都可以直接调用这些方法
	   增强了方法的集成性，并且避免了方法的
	*/
	// fmt.Println(slice.Plus())
	// fmt.Println(slice.Minus())
	// fmt.Println(slice.Mul())
	// fmt.Println(slice.Dev())

	// var studentMap StudentMap = map[string]interface{}{
	// 	"name": "张三",
	// 	"age":  18,
	// }
	//
	// studentMap.setName("李四")
	// studentMap.setAge(27)
	//
	// fmt.Println(studentMap)

	fmt.Println(slice.compute("PLUS"))
	fmt.Println(slice.compute("MINUS"))
	fmt.Println(slice.compute("MULTIPLY"))
	fmt.Println(slice.compute("DIVISION"))
}

// 给自定义类型提供方法
// 给自定义类型的所有数据提供方法
type ComputeSlice []int

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

// func (cs ComputeSlice) Plus() int {
// 	return cs[0] + cs[1]
// }
//
// func (cs ComputeSlice) Minus() int {
// 	return cs[0] - cs[1]
// }
//
// func (cs ComputeSlice) Mul() int {
// 	return cs[0] * cs[1]
// }
// func (cs ComputeSlice) Dev() int {
// 	return cs[0] / cs[1]
// }

// ---------------------------------------------

// func Plus(a int, b int) int {
// 	return a + b
// }
//
// func Minus(a int, b int) int {
// 	return a - b
// }
//
// func Mul(a int, b int) int {
// 	return a * b
// }
// func Dev(a int, b int) int {
// 	return a / b
// }

// ----------------------------------------------

// type StudentMap map[string]interface {
// }
//
// func (sm StudentMap) setName(name string) {
// 	// TODO
// 	sm["name"] = name
// }
//
// func (sm StudentMap) setAge(age int) {
// 	// TODO
// 	sm["age"] = age
// }
