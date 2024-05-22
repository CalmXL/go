package main

import "fmt"

func plus(a, b interface{}) interface{} {
	// 获取 a 的类型
	switch a.(type) {
	case int:
		_a, _ := a.(int)
		_b, _ := b.(int)
		return _a + _b
	case float64:
		_a, _ := a.(float64)
		_b, _ := b.(float64)
		return _a + _b
	case string:
		_a, _ := a.(string)
		_b, _ := b.(string)
		return _a + _b
	default:
		panic("类型不支持")
	}
}

func main() {
	
	// var a interface{} = 1
	// // 无效运算: a += 1(类型 interface{} 和 untyped int 不匹配)
	// // a += 1
	// // a.(断言类型)
	// a = a.(int) + 1func
	// fmt.Println(a)
	
	// ----------------------------------
	
	var a interface{} = 100
	res := plus(a, 1000)
	fmt.Println(res)
	
}
