package main

import "fmt"

func main() {
	// // 获取 slice 长度与容量
	// intSlice := []int{1, 2, 3}
	//
	// // 在未指定容量的前提下，容量与切片的长度相等
	// fmt.Printf("%d, %d", len(intSlice), cap(intSlice))
	//
	// // var newIntSlice []int
	// // fmt.Printf("%d, %d", len(newIntSlice), cap(newIntSlice)) // => 0, 0
	//
	// intSlice = append(intSlice, 4)
	// /**
	//   切片扩容
	// 	oldCap < 256
	// 		newCap => oldCap * 2
	// 	oldCap >= 256
	// 		newCap => oldCap + oldCap / 4 累加 192
	// */
	// fmt.Printf("%d, %d", len(intSlice), cap(intSlice)) // 4, 6

	// -------------------------------------------------------------
	/**
		type slice struct
	    array unsafe.Pointer  => init address
			len   int
			cap   int
	*/
	// intSlice := []int{1, 2, 3}
	// newIntSlice := append(intSlice, 4)
	//
	// fmt.Printf("%v, %p, %d, %d\r\n", intSlice, &intSlice[0], len(intSlice), cap(intSlice))
	// fmt.Printf("%v, %p, %d, %d\r\n", newIntSlice, &newIntSlice[0], len(newIntSlice), cap(newIntSlice))

	// intSlice := make([]int, 3, 5)
	//
	// intSlice[0] = 1
	// intSlice[1] = 2
	// intSlice[2] = 3
	//
	// // 如果 append 的过程中引起了扩容，那么将重新分配内存空间，存储元素，并但会新的切片对象
	// // 如果没有引起扩容，将返回原本的切片引用
	// newIntSlice := append(intSlice, 4)
	// fmt.Printf("%v, %p, %d, %d\r\n", intSlice, &intSlice[0], len(intSlice), cap(intSlice))
	// fmt.Printf("%v, %p, %d, %d\r\n", newIntSlice, &newIntSlice[0], len(newIntSlice), cap(newIntSlice))

	// ----------------------------------------------------------------
	var intSlice []int

	for i := 0; i < 5; i++ {
		intSlice = append(intSlice, i)
	}

	fmt.Println(intSlice)

	// intSlice2 := make([]int, 5)
	intSlice2 := make([]int, 0, 5) // 没有长度但是有容量，可以使用 append
	for i := 0; i < 5; i++ {
		// 追加元素
		intSlice2 = append(intSlice2, i)

		// 填充元素
		// intSlice[i] = i
	}
	fmt.Println(intSlice2) // => [0 0 0 0 0 0 1 2 3 4]

}
