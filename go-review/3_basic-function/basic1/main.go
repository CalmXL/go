package main

import "fmt"

// syntax error: non-declaration statement outside function body
// a := 1
func main() {
	a, b := 1, 2
	c, d := 3, 4
	
	// res1 := a + b
	// res2 := c + d
	res1 := plus(a, b)
	res2 := plus(c, d)
	fmt.Println(res1, res2)
	
	// ❌ 函数内部不能使用 函数声明
	// func test(){}
	test := func() {
		a := 3
		fmt.Println(a)
		fmt.Println(b)
	}
	
	// test()
	
	test1(test)
}

func plus(a, b int) int {
	return a + b
}

// 回调
func test1(cb func()) {
	cb()
}

// 闭包
func test2() func() {
	return func() {
		fmt.Println("Clourse")
	}
}
