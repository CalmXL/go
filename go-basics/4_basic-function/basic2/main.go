package main

import "fmt"

/**
  1. 函数参数是什么？
    参数是为了让函数成为一个数据接口
		接口接收一些值到函数内部参与函数的程序

	2. 函数为什么需要参数？
		函数如果直接获取外部作用域的变量
    到函数内部参与函数程序，函数在其他没有这些变量的地方
		执行就会出现错误
		所以尽量保证定义函数的参数，能更灵活
		更具有复用性的使用函数
		我们总是希望函数成为标准的输入输出 (纯函数)

	形式参数（形参）：函数定义时对函数调用时接收值的描述 （参数的占位符）
	实际参数（实参）：函数在调用时插入的实际的值
	返回值: 函数执行完毕后输出的值，可以在函数调用的时赋值给一个变量
*/

/**
 纯函数
	输入相同的值，总会有固定的值输出
*/
// 函数的接口
// 通过参数来表明函数执行时接收两个值，并且两个值都是 int
func plus(a int, b int) int {
	// 值的输出
	return a + b
}

func test1() {

}

// func test1() {
// 	num1 := 1
// 	num2 := 2
//
// 	plus := func() {
// 		fmt.Println(num1 + num2)
// 	}
// }
//
// func test2() {
// 	plus := func() {
// 		fmt.Println(num1 + num2)
// 	}
// }

// -------------------------------------
// 参数的写法
func getBill(pay, balance float32, productName string) string {
	return fmt.Sprintf(`
		付款：%.2f\r\n
		余额：%.2f\r\n
		商品名称：%s\r\n
	`, pay, balance, productName)
}

// ---------------------------------------
// 可变参数
// args => arguments
func computeSum(args ...int) int {
	// args => {[]int} len 5 cap 5
	sum := 0
	for _, num := range args {
		sum += num
	}
	return sum
}

// -----------------------------------
// 多返回值
func getSum(desc string, args ...int) (int, []int, string) {
	sum := 0
	for _, num := range args {
		sum += num
	}
	return sum, args, desc
}

// ------------------------------------------------------
// 返回值类型指定返回变量名

func computeSum3(args ...int) (sum int, argsSlice []int) {
	for _, num := range args {
		sum += num
	}
	argsSlice = args

	return sum, argsSlice
}

func main() {
	// res := plus(1, 2)
	// fmt.Println(res)

	// test1() 用作值，但它不返回任何内容
	// res2 := test1()

	// res := getBill(32.5, 22.5, "水杯")
	// fmt.Println(res)

	// res := computeSum(1, 2, 3, 4, 5)
	// fmt.Println(res)

	// sum, args, desc := getSum("求和", 1, 2, 3)
	// fmt.Println(sum, args, desc)

	sum, args := computeSum3(1, 2, 3, 4)
	fmt.Println(sum, args)
}
