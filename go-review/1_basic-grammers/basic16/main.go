package main

import "fmt"

func main() {
	
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }
	//
	// i := 0
	// for ; i < 3; i++ {
	// 	fmt.Println(i)
	// }
	//
	// j := 0
	// for ; j < 3; {
	// 	fmt.Println(j)
	// 	j++
	// }
	
	// i := 0
	// for {
	// 	if i < 10 {
	// 		fmt.Println(i)
	// 		i++
	// 	} else {
	// 		break
	// 	}
	// }
	
	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("偶数: %d\r\n", i)
	// 	} else {
	// 		fmt.Printf("奇数: %d\r\n", i)
	// 	}
	// }
	
	// 	for i := 0; i < 100; i++ {
	// 		if i%2 == 0 {
	// 			fmt.Printf("偶数: %d\r\n", i)
	// 			// 跳出本次循环
	// 			continue
	// 		}
	//
	// 		fmt.Printf("奇数: %d\r\n", i)
	// 	}
	
	// 求 1 -100 的 和
	// sum := 0
	// for i := 0; i <= 100; i++ {
	// 	sum += i
	// }
	//
	// fmt.Println(sum)
	
	// 分别求奇数与偶数的和
	// oddSum, evenSum := 0, 0
	//
	// for i := 0; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		evenSum += i
	// 	} else {
	// 		oddSum += i
	// 	}
	// }
	//
	// fmt.Println(oddSum, evenSum)
	
	// 求 2 - 100 以内的质数的和
	sum := 0
	isPrimNumber := true
	for i := 2; i < 100; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrimNumber = false
				break
			}
		}
		if isPrimNumber {
			fmt.Println(i)
			sum += i
		}
		isPrimNumber = true
	}
	
	fmt.Println(sum)
}
