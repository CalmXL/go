package main

import "fmt"

/*
*

	  Collection 集合
			装载很多值的容器
				a := 1
				b := 2
				c := 3
				[ a, b, c ] => 集合

		集合的类型：
			array 数组
			map   字典、哈希表、映射
			slice 切片
			list  列表

		数组: 编程语言当中，最常见的有序列表的数据结构

		1. [ 1, 2, 3, 4, 5 ]
		2. [ 5, 4, 3, 2, 1 ]
		总结:
			1.装载的值一致
			2. 值的顺序不一致(有序地列表，顺序不一致，数据含义就不同)

		无序列表：key 和 key 不存在顺序关联
		{
			name: "小野森森",
			age: 37
	  }

		数组的基本形式: [1, 2, 3, 4, 5]
						索引:  0, 1, 2, 3, 4
			元素（element）: 数组内部存储的值
			索引: 表示元素在数组中的位置

			numbers => [1, 2, 3, 4, 5]
					number[2] => 3
					number[2] => 30

			数组的特点:
				1. 数组的有序性：由索引来描述元素在数组中的位置
				2. 数组的可遍历性：通过有序性的特征，可以将元素按照索引顺序进行迭代
				3. 数组的定长性: 数组在声明定义时，就必须指定长度
*/
func main() {
	// 数组的声明
	/**
	  数组的类型: 长度 + 元素类型
		string => 数组可以存放什么类型的数组
		[] => classOne 是数组或者切片类型
	  [3] => classOne 是一个长度为 3 的数组。

	  java 中
		String [] classOne = new String[3]
		Go 中
		var classOne [3]string
	*/
	// var classOne [3]string

	// var classTwo [4]string
	// var classThree [3]string

	// 数组元素的赋值
	// classOne[0] = "张三"
	// classOne[1] = "李四"
	// classOne[2] = "王五"

	// 数组元素的修改
	// classOne[2] = "赵六"

	// 数组不能增加元素
	// 无效的 数组 索引 '3' (3 元素的数组超出界限)
	// classOne[3] = "赵六"

	// classThree[0] = "李四"
	// classThree[1] = "张三"
	// classThree[2] = "王五"

	// fmt.Printf("%T, %T\r\n", classOne, classTwo) // [3]string [4]string

	/**
	  数组的相等性判断
			1. 数组的元素类型要一致
			2. 数组的元素长度要一致
			3. 数组的元素要一一对应，值一致
	*/
	// 说明值相等
	// fmt.Println(classOne == classThree) // => false

	/**
	  只要元素类型与元素的长度一致，就是一个类型
	*/

	// 类型一致
	// classOne = classThree
	// fmt.Println(classOne)
	// fmt.Println(classThree)

	// 数组的遍历
	// for i := 0; i < len(classOne); i++ {
	// 	fmt.Println(classOne[i])
	// }

	// for index, el := range classOne {
	// 	fmt.Println(index, el)
	// }

	// 默认值
	// var nums [3]int
	// fmt.Println(nums) // => [0 0 0]

	/**
	  数组的初始化
		1. var classOne [3]stirng
		2. classOne := [3]string{"张三", "李四", "王五"}
		3. classOne := [3]string{1: "张三", 2: "李四"}


	*/

	// classOne := [3]string{"张三", "李四", "王五"}
	// fmt.Println(classOne)

	// ----------------------------------------------
	// classOne := [3]string{1: "张三", 2: "李四"}
	//
	// for index, el := range classOne {
	// 	fmt.Println(index, el)
	// }
	// 0
	// 1 张三
	// 2 李四

	// ----------------------------------------------
	classOne := [...]string{"张三", "李四", "王五"}
	// 但是依旧不能扩容
	// 无效的 数组 索引 '3' (3 元素的数组超出界限)
	// classOne[3] = "赵六"
	for index, el := range classOne {
		fmt.Println(index, el)
	}

}
