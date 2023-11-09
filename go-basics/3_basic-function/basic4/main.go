package main

/**
  同步和异步机制

	回调函数(callback) => 参数传递的值是一个函数

	回调函数的作用：
		把函数任务分步进行处理
		函数执行期间，一部分任务交给回调函数进行处理
*/

// func step1(a int, b int) int {
// 	return a + b
// }
//
// func step2(a int, b int) int {
// 	return a - b
// }
//
// func step3(res int) {
// 	fmt.Println(res)
// }
//
// func step4(res int) {
// 	fmt.Println(res)
// }

// func test(cb func()) {
// 	cb()
// }

func test(strFn func(str string) string) string {
	str := "abc"
	finalStr := strFn(str)

	return finalStr
}

func main() {
	// a := 1
	// b := 2

	// 代码的执行过程就是同步执行
	// res1 := a + b
	// res2 := a - b
	//
	// fmt.Println(res1)
	// fmt.Println(res2)

	// 完成一个任务，在执行下一个任务 => 同步执行任务
	// res1 := step1(a, b)
	// res2 := step2(a, b)
	// step3(res1)
	// step4(res2)

	/**
	  执行一个任务，未完成任务的前提下，执行下一个任务 => 异步执行任务
		执行顺序：
			1. step1(未完成执行)
	    2. step2(未完成执行)
			3. step3

		res1 := step1(a, b)
		res2 := step2(a, b)
		step3(res1)
		step4(res2)
	*/
}
