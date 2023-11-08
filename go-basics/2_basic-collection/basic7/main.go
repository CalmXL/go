package main

import "fmt"

func main() {
	intSlice := make([]int, 3, 5) // => [0, 0, 0]

	// 填充元素
	intSlice[0] = 1
	intSlice[1] = 2
	intSlice[2] = 3

	// 追加元素 append
	/**
	  [1, 2, 3, 4], 容量充足，不需要扩容
		len++ => 4 < 5 => 不需要扩容

	  1. 扩容，新开辟一个 slice 的连续空间
		2. 将原有空间的存的元素移动到新的空间
		3. 通过 append 返回一个新的 slice 的引用

		append 在不扩容的情况下，返回原有的 slice 引用
					 在扩容的情况下，返回新的 slice 引用

		扩容选择新开辟连续空间的原因？
			原有的连续空间不能保证后续的预扩容空间内没有其他值的存储
			所以 Go 选择重新开辟一个预容量的连续空间的目的
	*/
	intSlice = append(intSlice, 4)

	// newSlice := intSlice[1:4]

	// append(int[]{1, 2}, 4, 5, 6) => newSlice
	newSlice := append(intSlice[:2], intSlice[3:]...)

	fmt.Println(newSlice)
}
