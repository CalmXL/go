package main

import "fmt"

/**
  循环： Loop
	计步式循环（计数式）
三要素:
	步数、循环条件、步数的操作
	0     <=10   +1
for(i := 0; i <= 10; i++) {}
i: 迭代变量 => 初始化为 0
每次循环，1. 检查 i <= 10
				2. 执行循环程序
				3. 步数的操作 (迭代操作)

for (迭代变量初始化；循环条件；迭代操作){
	// 循环体 / 循环代码块 / 循环程序
}

Go 语言中，没有 while, 也没有 do while
for => 代替 => while
*/

func main() {
	/**
	  1. var i = 0
		2. i < 3
		3. fmt.Println(0)
		4. i ++

		2. 1 < 3
		3. fmt.Println(1)
	  4. i ++

		2. 2 > 3
		3. fmt.Println(2)
	  4. i ++

		2. 3 < 3 x => 循环结束
	*/

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	// ----------------------------------

	// i := 0
	// for ; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	// ----------------------------------

	// i := 0
	// for ; i < 3; {
	// 	fmt.Println(i)
	// 	i++
	// }

	// ------------------------------------
	// while (i < 10) {
	// 	fmt.Println(i)
	// 	i ++
	// }

	// i := 0
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// i := 0
	// for { // while(true)
	//
	// 	// if i < 10 {
	// 	// 	fmt.Println(i)
	// 	// 	i++
	// 	// } else {
	// 	// 	// 断开 => 跳出循环 => 结束循环
	// 	// 	break
	// 	// }
	//
	// 	// ------------------------------------
	//
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	/**
	  得到100以内的奇数和偶数
	*/
	/**
	  1. 3 % 2 == 1
		2. continue
		3. 3 <= 100
		4. i++
		5. 4 % 2 != 1
		6. fmt.Println(4)
	*/
	// for i := 0; i <= 100; i++ {
	// 	if i%2 == 1 {
	// 		// fmt.Println(i)
	// 		// 跳出本次循环
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	/**
	  求 1 - 100 的总和
	*/
	// sum := 0
	// for i := 0; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	/**
	  求奇偶的总和
	*/
	// 	oddSum := 0
	// 	evenSum := 0
	//
	// 	for i := 0; i <= 100; i++ {
	// 		switch i % 2 {
	// 		case 0:
	// 			evenSum += i
	// 		case 1:
	// 			oddSum += i
	// 		default:
	// 			break
	// 		}
	// 	}
	// 	fmt.Println(oddSum, evenSum)
	// }

	/**
	  求 2 - 100 内的质数总和
	*/
	primNumberSum := 0
	isPrimNumber := true

	for i := 2; i <= 100; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrimNumber = false
				break
			}
		}
		if isPrimNumber {
			fmt.Println(i)
			primNumberSum += i
		}

		isPrimNumber = true
	}

	fmt.Println(primNumberSum)
}
