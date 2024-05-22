package main

import "fmt"

// func plus(a int, b int) int { return a + b }
//
// func test1() {
// 	num1 := 1
// 	num2 := 2
//
// 	plus := func() {
// 		fmt.Println(num1 + num2)
// 	}
//
// 	plus()
// }

// func getBill(pay, balance float32, productName string) string {
// 	return fmt.Sprintf(`
// 	付款: %.2f
// 	余额：%.2f
// 	商品名称：%s
// 	`, pay, balance, productName)
// }

// func computeSum(args ...int) int {
// 	sum := 0
// 	for _, num := range args {
// 		sum += num
// 	}
// 	return sum
// }

// func getSum(desc string, args ...int) (int, []int, string) {
// 	sum := 0
// 	for _, num := range args {
// 		sum += num
// 	}
//
// 	return sum, args, desc
// }

func computeSum(args ...int) (sum int, argsSlice []int) {
	for _, num := range args {
		sum += num
	}
	
	argsSlice = args
	
	return sum, argsSlice
}

func main() {
	// res := plus(1, 2)
	// // test1() 用作值，但它不返回任何内容
	// res2 := test1()
	//
	// fmt.Println(res)
	// fmt.Println(res2)
	
	// res := getBill(32.5, 22.5, "水杯")
	// fmt.Println(res)
	
	// res := computeSum(1, 2, 3, 4, 5)
	// fmt.Println(res)
	
	// sum, args, desc := getSum("求和", 1, 2, 3)
	// fmt.Println(sum, args, desc)
	
	sum, args := computeSum(1, 2, 3, 4)
	fmt.Println(sum, args)
}
