package main

import "fmt"

func main() {
	// 声明一个数组长度为 3 的字符串数组
	// var studentArr [3]string
	//
	// studentArr[0] = "张三"
	// studentArr[1] = "李四"
	// studentArr[2] = "王五"
	//
	// // invalid array index 3 (out of bounds for 3-element array)
	// // 数组是靠长度来规范其是否能添加元素的
	// studentArr[3] = "赵六"

	// =--------------------------------------
	// var studentArr []string

	// runtime error: index out of range [0] with length 0
	// 编译通过，运行错误
	// 数组未指定长度，默认为 0
	// studentArr[0] = "张三"

	// --------------------------------------------
	/**
	  声明一个切片就是声明一个默认长度为 0 的数组 （不能在中括号内部写任何数字）
		切片的本质就是动态数组 （数组的长度是自动扩容的）
		在声明切片的时候，长度是不可以设置的，如果设置了长度的声明就会声明一个数组出来
		切片的底层是基于数组实现的
	*/
	// var studentArr []string

	// 数组必须要先分配空间，再存储元素
	// 在已有的数组空间中存储元素
	// studentArr[0] = "张三"

	/**
	  一个长度为 0 的数组需要使用 append 方法进行元素的添加
		append 是全局方法，并返回新的引用，并且必须赋值给一个变量
		切片(slice)需要使用 append 方法进行元素的追加
	*/
	// 在存储之前，长度扩容后， 在进行存储
	// studentArr = append(studentArr, "张三", "李四")

	// -------------------------------------------
	// var studentArr1 []string
	// var studentArr2 []string
	//
	// // 无效运算: studentArr1 == studentArr2 (在 []string 中未定义运算符 ==)
	// // invalid operation: studentArr1 == studentArr2 (slice can only be compared to nil)
	// // 无效     操作                                    切片 只能跟 nil 进行比较
	// fmt.Println(studentArr1 == studentArr2)

	// ----------------------------------------------

	// var studentArr1 []string
	// var studentArr2 []string

	/**
	  append 会返回一个新的切片
	  新的切片必须赋值给原切片变量，才能做到增加元素
	*/

	// studentArr2 = append(studentArr1, "张三", "李四", "王五")
	// fmt.Println(studentArr1, studentArr2) // [] [张三 李四 王五]
	// *****************************
	// studentArr1 = append(studentArr1, "张三", "李四", "王五")
	// studentArr2 = append(studentArr1, "张三", "李四", "王五")
	// fmt.Println(studentArr1, studentArr2) // => [张三 李四 王五] [张三 李四 王五 张三 李四 王五]

	// -----------------------------------------------------------------
	/**
	  切片（slice）的初始化
			1. 默认长度为 0 的数组
				var studentArr []string
					studentArr = append(studentArr, "xx", "xx" ...)
			2. 定义切片的初始化
					studentArr := []string{"张三", "李四", "王五"}
			3. make 方法初始化存储空间
					studentArr := make([]string, 2)
	*/

	// studentArr := []string{"张三", "李四", "王五"}

	// 通过 make 分配切片的存储空间来指定默认的长度
	// 好处: 由于切片需要扩充空间
	//			如果先知道需要存储多少元素，那么就先指定多少空间
	//      扩容的次数变少了
	studentArr := make([]string, 2)

	studentArr[0] = "张三"
	studentArr[1] = "李四"
	// runtime error: index out of range [2] with length 2
	// studentArr[2] = "王五"

	studentArr = append(studentArr, "王五")

	fmt.Println(studentArr) // => [张三 李四 王五]

}
