package main

import "fmt"

// func test() func() {
// 	return func() {}
// }
//
// func test1(c int, d int) func(a int, b int) {
// 	return func(a int, b int) {}
// }
//
// func test2(a int, b int) func(aa int, bb int) func(aaa int, bbb int) {
// 	return func(a int, b int) func(aa int, bb int) {
// 		return func(a int, b int) {}
// 	}
// }
//
// func test3(cb1 func(a int), cb2 func(b int)) func(aa int, bb int) func(aaa int, bbb int) {
// 	return func(aa int, bb int) func(aaa int, bbb int) {
// 		cb1(1)
// 		return func(aaa int, bbb int) {
// 			cb2(2)
// 		}
// 	}
// }

// 闭包
// var a = 1
//
// func test() {
// 	b := 2
// 	test1 := func() {
// 		// 闭包的函数
// 		// test1 环境捆绑了 test 的环境
// 		// 因此 test1 作用域可以访问 test 作用域内部的变量
// 		// test1 不仅捆绑了 test 的环境，同时也捆绑了全局作用域的环境
//
// 		/**
// 		    test1 scope =>  test1   ev
// 												test    ev
// 												global  ev
// 		*/
// 		fmt.Println(a, b)
// 	}
//
// 	test1()
// }

// 未逃逸
func test() {
	a := 1
	
	test1 := func() {
		fmt.Println(a)
	}
	
	test1()
}

func test2() func() {
	a := 1
	test1 := func() {
		fmt.Println(a)
	}
	
	return test1
}

func compute(count int) (func(num int) int, func(num int) int) {
	increase := func(num int) int {
		count += num
		return count
	}
	
	decrease := func(num int) int {
		count += num
		return num
	}
	
	return increase, decrease
}

func main() {
	
	// fn := test2()
	// fn()
	//
	
	increase, decrease := compute(2)
	
	res := increase(3)
	res2 := decrease(3)
	
	fmt.Println(res, res2)
}
