package main

import "fmt"

/*
	类型转换: 类型转换都是为了运算
	类型断言: 一个不明确的类型（interface{}）转换为一个明确的类型
*/

// func plus(a int, b int) int {
// 	return a + b
// }
//
// func plus[T string | int | float64](a T, b T) T {
// 	return a + b
// }

// func plus(a, b interface{}) interface{} {
// 可以获得 a 的类型
// 	switch a.(type) {
// 	case int:
// 		_a, _ := a.(int)
// 		_b, _ := b.(int)
// 		return _a + _b
// 	case float64:
// 		_a, _ := a.(float64)
// 		_b, _ := b.(float64)
// 		return _a + _b
// 	case string:
// 		_a, _ := a.(string)
// 		_b, _ := b.(string)
// 		return _a + _b
// 	default:
// 		panic("类型不支持")
// 	}
// }

type Todo struct {
	id        int
	content   string
	completed bool
}

type TodoMap map[string]any

func main() {

	// 接口类型的变量.(断言的类型)
	// var a interface{} = 1
	// // 无效运算: a += 1(类型 interface{} 和 untyped int 不匹配)
	// // a += 1
	// a = a.(int) + 1
	// fmt.Println(a)

	// -----------------------------
	// var a int = 1
	// var b float64 = 2
	// // 无效的类型断言: b.(int) (左侧为非接口类型 float64)
	// // res := a + b.(int)
	// res := a + int(b)
	// fmt.Println(res) // => 3

	// ------------------------------
	// res1 := plus(1, 2)
	// res2 := plus(1.1, 2.1)
	// res3 := plus("a", "b")
	// res4 := plus([]int{1, 2, 3}, "b")
	//
	// fmt.Println(res1)
	// fmt.Println(res2)
	// fmt.Println(res3)
	// fmt.Println(res4)

	todoList1 := make([]Todo, 0)
	todoList2 := make([]TodoMap, 0)

	setTodo := func(todo any) {
		switch todo.(type) {
		case Todo:
			todoList1 = append(todoList1, todo.(Todo))
		case TodoMap:
			todoList2 = append(todoList2, todo.(TodoMap))
		default:
			panic("todo 类型错误")
		}
	}

	setTodo(Todo{
		id:        1,
		content:   "Todo",
		completed: false,
	})

	setTodo(TodoMap{
		"id":        1,
		"content":   "Todo",
		"completed": false,
	})

	fmt.Println(todoList1)
	fmt.Println(todoList2)
}
