package main

import "fmt"

/*
*

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
	// myInfo := make(map[string]string, 3)
	//
	// myInfo["name"] = "xiaoyesensen"
	// myInfo["age"] = "18"
	//
	// fmt.Println(myInfo)
	// fmt.Println(len(myInfo)) // => 2
	//
	// myInfo["height"] = "176"
	// // 从超过容量的 key 开始，追加需要扩容
	// myInfo["hobby"] = "music"
	//
	// fmt.Println(myInfo)

	// -------------------------------------------
	// 重复 key
	// myInfo := map[string]string{
	// 	"name": "小野森森",
	// 	"age":  "18",
	// }
	// // 修改操作
	// // key 在集合当中是唯一的
	// myInfo["age"] = "20"
	// fmt.Println(myInfo)

	// -------------------------------------------
	// 读取、修改
	// stringMap := make(map[string]string, 1)
	// stringMap["name"] = "xiaotyesensen"
	// stringMap["name"] = "superxiaotye"
	//
	// fmt.Println(stringMap["name"])

	// intMap := make(map[int]string, 3)
	//
	// intMap[1001] = "张三"
	// intMap[1002] = "李四"
	// intMap[1003] = "王五"
	//
	// fmt.Println(intMap)

	// boolMap := make(map[bool]int, 2)
	//
	// boolMap[true] = 1
	// boolMap[false] = 0
	//
	// fmt.Println(boolMap)

	// mySlice := []string{"张三", "李四", "王五"}
	//
	// myMap := make(map[string][]string, 2)
	//
	// myMap["classOne"] = mySlice
	//
	// fmt.Println(myMap)

	// myMap := make(map[string]string, 3)
	//
	// myMap["name"] = "xiaoyesensen"
	// myMap["age"] = "18"
	// myMap["height"] = "176"
	// fmt.Println(myMap["weight"]) // => 返回空字符串
	//
	// _, exist := myMap["weight"]
	// fmt.Println(exist) // '' false
	//
	// _, exist2 := myMap["height"]
	// fmt.Println(exist2) // true

	// myMap := make(map[int]int, 3)
	//
	// myMap[0] = 1
	// myMap[1] = 2
	// myMap[2] = 3
	// fmt.Println(myMap[3]) // => 0

	// ----------------------------------------------------
	// 写入

	// var mySlice [] string
	// 只有声明没有初始化
	// 初始化 var myMap map[int]int {}
	// var myMap map[int]int
	// assignment to entry in nil map
	// 没有进行初始化
	// myMap[0] = 1
	// fmt.Println(myMap)

	// stringMap := make(map[string]string, 1)
	// stringMap["name"] = "xiaoyesensen"

	// ----------------------------------------------------
	// 删除 key
	// stringMap := make(map[string]string, 2)
	// stringMap["name"] = "xiaoyesensen"
	// stringMap["age"] = "18"
	//
	// fmt.Println(stringMap)
	//
	// // delete 内置方法  =>   stringMap.delete(element)
	// delete(stringMap, "age")
	// fmt.Println(stringMap)
	//
	// // 删除不存在的 key, 不会有任何操作和报错
	// delete(stringMap, "height")

	// -----------------------------------------------------

	myInfo := make(map[string]string, 3)
	myInfo["name"] = "xulei"
	myInfo["age"] = "27"
	myInfo["height"] = "178"
	myInfo["hobby"] = "music"

	// for range 针对于 map 的 key 的枚举是无序的
	// 枚举的过程是无序且顺序随机的
	for key, value := range myInfo {
		fmt.Println(key, value)
	}
}
