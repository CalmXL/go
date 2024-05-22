package main

import "fmt"

/*
*

		迭代:0 - 1 - 2 - 3
			对于一个数据集合，每一次按照顺序获取一个数据的过程叫做迭代
	  遍历: 0 - 3
			对于一个数据集合，所有迭代的过程叫做遍历

		迭代是可以通过遍历实现的
		所有迭代的过程加在一起就是遍历的过程

		[ 1, 2, 3, 4, 5 ]
			0  1  2  3  4

		字符换
		1. 有序性 hi => ih
	  2. 集合性 [ 'h', 'i' ]

		for range => 遍历 => 有序集合使用: 字符串、数组、切片、map、channel
		for       => 循环

		for key, value := range var {

		}

		key, value ?
		无序列表
		{
			key   value
			name: "小野森森",
			age: 37
		}


		有序列表
							      [1, 2, 3, 4, 5]
		index: 索引       0, 1, 2, 3, 4
		element: 列表元素  1, 2, 3, 4, 5

	  {
	    0 : 1,
			1 : 2,
			2 : 3,
			3 : 4,
			4 : 5
	  }
*/
func main() {
	// name := "xiaoyesensne-小野森森"
	// nameRune := []rune(name)
	//
	// for index, element := range name {
	// 	// element 从字符串遍历出来时 ASCII 码
	// 	// fmt.Println(index, element)
	// 	// fmt.Println
	// 	fmt.Printf("%d, %c\r\n", index, element)
	// 	// fmt.Printf("%d, %c\r\n", index, name[index])
	// 	element = 211
	// }
	//
	// fmt.Println(name)
	//
	// for index, _ := range nameRune {
	// 	fmt.Printf("%d, %c\r\n", index, nameRune[index])
	// }

	// --------------------------------------------------
	username := "xiaoyesensen"
	password := "xyss123456"

	for _, pc := range password {
		for _, uc := range username {
			if pc == uc {
				fmt.Println("有重复字母")
				return
				// goto hasRepeatChar
			}
		}
	}

	// hasRepeatChar:
	// 	fmt.Println("有重复字母")
}
