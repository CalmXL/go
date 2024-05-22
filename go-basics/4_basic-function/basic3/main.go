package main

import "fmt"

/**
  函数类型:
    基本的形式：func ()
    带参的形式: func (a int, b int)
		带返回值的形式: func () int | func() func()

  执行符号
		test()
*/

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
// 		return func(aaa int, bbb int) {
//
// 		}
// 	}
// }
//
// func test3(cb1 func(a int), cb2 func(b int)) func(aa int, bb int) func(aaa int, bbb int) {
// 	return func(a int, b int) func(aa int, bb int) {
// 		cb1(1)
// 		return func(aaa int, bbb int) {
// 			cb2(2)
// 		}
// 	}
// }

// ---------------------------------------------

/**
  闭包 ===> 封闭的作用域空间
			主要针对函数内部的一个名词
			函数内部的作用域捆绑这外部的作用域环境，这种现象叫做闭包

	闭包特性：
		1. 内部函数一直捆绑着外部作用域环境

*/
// var a = 1
//
// func test() {
// 	b := 2
// 	test1 := func() { // 闭包的函数
// 		fmt.Println(a, b)
// 		// test1 环境捆绑了 test 的环境
// 		// 所以 test1 作用域可以访问 test 作用域内部的变量
// 		// test1 不仅捆绑了 test 的环境，同时也捆绑了全局作用域的环境
// 		/**
// 				test1 scope => test1  ev
// 		                   test   ev
// 		                   global ev
// 		*/
// 	}
// 	test1()
// }

// ---------------------------------------
/**
  闭包特性：
		1. 内部函数可以访问到外部环境的变量
		2. 内部函数可以访问到外部函数的参数
		3. 内部函数可以操作外部环境的变量和函数参数
		4. 闭包函数可以传入参数进行运算
		5. 闭包使外部函数的变量或者参数，成为内部函数的私有化变量

	闭包的好处：
		1. 延长局部变量的生命周期
		2. 形成类似于面相对象的外部私有化
		3. 使外部作用域可访问到内部作用域的变量

	逃逸分析：
    Go 编译器的优化技术
		要确定一个内部变量是否要在堆内存上分配空间的一种分析技术

    // 未逃逸
		func test() {
      a := 1
			test1 := func () {
					Println(a)
			}

			test1()
		}

		// 逃逸 => 逃逸分析
		func test() {
      a := 1
			test1 := func () {
					Println(a)
					// 当内部作用域访问了外部作用域的变量，这个变量就是持久化变量
					// 闭包会触发变量逃逸分析
					// 1. 内部函数访问了外部变量 => 分配堆内存空间给这个变量
          // 2. 内部函数没有访问外部变量 => 不分配堆内存空间给外部变量
			}

      return test1
		}
		test1 := test()
*/
// func test() func() {
// 	count := 0
//
// 	return func() {
// 		fmt.Println(count)
// 	}
// }

// func test(count int) func() {
//
// 	return func() {
// 		fmt.Println(count)
// 	}
// }

// func test(count int) func() int {
// 	return func() int {
// 		count += 1
// 		return count
// 	}
// }

// func test(count int) func(num int) int {
// 	return func(num int) int {
// 		count += num
// 		return count
// 	}
// }

func test(count int) (func(num int) int, func(num int) int) {
	increase := func(num int) int {
		count += num
		return count
	}

	decrease := func(num int) int {
		count -= num
		return count
	}

	return increase, decrease
}

func main() {
	// test()

	// test1 := test(1)
	// test1()

	// test1 := test(1)
	// count := test1()
	// fmt.Println(count) //

	// test1 := test(1)
	// count := test1(100)
	// count2 := test1(2)
	// fmt.Println(count, count2)

	increase, decerease := test(100)

	/**
	  101 103 106 103 101 100
	*/

	res1 := increase(1)
	res2 := increase(2)
	res3 := increase(3)

	res4 := decerease(3)
	res5 := decerease(2)
	res6 := decerease(1)

	fmt.Println(res1, res2, res3, res4, res5, res6)

}
