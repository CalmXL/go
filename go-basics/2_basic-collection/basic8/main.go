package main

import "fmt"

/**
  map类型：字典类型、映射

	英文词典 => m 开头的单词 => map => map 对应的释义
	map: 英文词典中的一个 key (键、键名、属性名、索引)
	map的释义: 英文词典中一个单词 key 对应的 value(键值，属性值)
	key-value: 键值对

	英文词典：{
		map: "词典",
		mathematics: "数学",
		matrix: "矩阵"
  }

	map 定义: 是一种包含多个 key-value 键值对的集合
		{
			name: "xiaoyesensen",
			age: "18",
		}

	map 的数据特点：
		1. map 是无序列表集合
		2. map 集合装载 key => value 的结构
		3. 一个 key 在 map 中具有唯一性
		4. 读、写、删除效率比较高
		5. slice、map、func 不能作为 key
			key 在系统中插入中要对比重复性，不能作为比较运算的类型不能做 key
*/
func test1() {}
func test2() {}

func main() {
	// 值的相等性
	// var slice1 []int
	// var slice2 []int
	//
	// fmt.Println(slice1 == slice2)

	// map1 := map[string]string{}
	// map2 := map[string]string{}
	// 无效运算: map1 == map2 (在 map[string]string 中未定义运算符 ==)
	// fmt.Println(map1 == map2)

	// fmt.Println(test1 == test2) // ❌

	// --------------------------------------------------------
	// 声明 map 并初始化
	// [key type] value type
	// myInfo := map[string]string{}
	// myInfo["name"] = "xiaoyesensen"
	// myInfo["age"] = "18"
	//
	// fmt.Println(myInfo)

	// 2.
	// 注意最后一个 key-value 后面必须打逗号
	// nameKey := "name"
	// myInfo := map[string]string{
	// 	// "name":  "xiaoyesensen",
	// 	nameKey: "xulei",
	// 	"age":   "18",
	// }
	//
	// fmt.Println(myInfo)

	// 3.
	/**
	    参数1： map 类型的数据
			参数2： 容量可选
	*/
	myInfo := make(map[string]string, 3)

	myInfo["name"] = "xiaoyesensen"
	myInfo["age"] = "18"

	fmt.Println(myInfo)
	fmt.Println(len(myInfo)) // => 2

	myInfo["height"] = "176"
	// 从超过容量的 key 开始，追加需要扩容
	myInfo["hobby"] = "music"

	fmt.Println(myInfo)
}
