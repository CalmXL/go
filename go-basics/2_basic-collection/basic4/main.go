package main

import "fmt"

/**

 */
func main() {
	/**
	  1. 访问元素
	*/
	// studentSlice := []string{"张三", "李四", "王五"}
	//
	// fmt.Println(studentSlice[1]) // => 李四
	//
	// studentSlice[1] = "xulei"
	// fmt.Println(studentSlice[1]) // => xulei

	/**
	  2. 截取 [start:end] => [start, end)
			1. [start:end]
			2. [1: len(studentSlice)]
			3. [1:]
			4. [:4]
			5. [:]
	*/

	// studentSlice := []string{"张三", "李四", "王五", "赵六", "小明", "小红"}
	// // studentSlice1 := studentSlice[1:4]  // => [李四 王五 赵六]
	// // studentSlice1 := studentSlice[0:len(studentSlice)] // => [张三 李四 王五 赵六 小明 小红]
	// // studentSlice1 := studentSlice[1:] // => [李四 王五 赵六 小明 小红]
	// // studentSlice1 := studentSlice[:4] // => [张三 李四 王五 赵六]
	// studentSlice1 := studentSlice[:]
	//
	// studentSlice[1] = "小李" // 会影响 studentSlice
	//
	// fmt.Println(studentSlice1)

	// ---------------------------------------
	/**
	  遍历
	*/
	// studentSlice := []string{"张三", "李四", "王五", "赵六", "小明", "小红"}
	//
	// for i := 0; i < len(studentSlice); i++ {
	// 	fmt.Println(studentSlice[i])
	// }
	//
	// for index, el := range studentSlice {
	// 	fmt.Println(index, el)
	// }

	// -----------------------------------------------------------------

	/**
	  添加元素
	*/

	// var classSlice []string

	/**
		func append(slice []Type, elems ...Type) []Type
	  参数1： 切片
	  参数2: string 类型的元素s
	*/
	// classSlice = append(classSlice, "张三", "李四", "王五")
	// studentSlice := make([]string, 3)
	//
	// studentSlice[0] = "小红"
	// studentSlice[1] = "小明"
	// studentSlice[2] = "小李"

	// for i := 0; i < len(studentSlice); i++ {
	// 	classSlice = append(classSlice, studentSlice[i])
	// }
	//
	// fmt.Println(classSlice)

	// 省略号操作
	// classSlice = append(classSlice, studentSlice...)
	// classSlice = append(classSlice, studentSlice[1:]...)
	// fmt.Println(classSlice)

	// -----------------------------------------------

	/**
	  删除元素 => 排除元素
	*/
	// studentSlice := []string{"张三", "李四", "王五", "赵六", "小明", "小红"}

	// 删除第一个元素
	// studentSlice = studentSlice[1:]
	// fmt.Println(studentSlice)

	// 删除最后一个元素
	// studentSlice = studentSlice[:len(studentSlice)-1]
	// fmt.Println(studentSlice)

	// 删除中间元素
	// studentSlice = append(studentSlice[:2], studentSlice[4:]...)
	// fmt.Println(studentSlice)

	// 空切片
	// studentSlice = []string{}
	// fmt.Println(studentSlice)

	// ----------------------------------------
	// 复制、拷贝
	classSlice := []string{"张三", "李四", "王五", "赵六", "小明", "小红"}

	// 两个切片变量指向同一个切片引用(引用赋值)
	// newClassSlice := classSlice
	//
	// classSlice[0] = "小野森森"
	//
	// fmt.Println(newClassSlice)
	//
	// newClassSlice[1] = "xulei"
	// fmt.Println(classSlice)

	// 拷贝
	// var newClassSlice []string
	newClassSlice := make([]string, len(classSlice))
	/**
		拷贝与被拷贝的切片是两个不同的切片

	  参数1：新的切片变量
		参数2: 源切片(source)
		copy 是不会扩充切片长度

		返回值: 拷贝切片元素的长度
	*/
	res := copy(newClassSlice, classSlice)
	fmt.Println(res, newClassSlice)

	newClassSlice[1] = "迪丽热巴"

	fmt.Println(newClassSlice, classSlice)
}
