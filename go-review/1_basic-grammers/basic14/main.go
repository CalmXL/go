package main

import "fmt"

func main() {
	// x := 1
	//
	// // 非布尔值 'x' (类型 int) 用作条件
	// if x {
	// 	fmt.Println("success")
	// }
	
	// x := 1
	// fmt.Scanln(&x)
	// fmt.Println(x)
	
	year, month := 0, 0
	
	fmt.Println("请输入年份")
	fmt.Scanln(&year)
	
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	
	if month == 2 {
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			fmt.Println("天数： 29")
		} else {
			fmt.Println("天数: 28")
		}
	} else {
		if month <= 12 {
			if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
				fmt.Println("天数: 31")
			} else {
				fmt.Println("天数: 30")
			}
		}
	}
}
