package main

func main() {
	/**
	  key 类型的问题
	*/
	// courseMap := make(map[string]string, 3)
	//
	// courseMap["go"] = "Golang课程"
	// courseMap["java"] = "Java WEB课程"
	// courseMap["rust"] = "Rust课程"
	//
	// fmt.Println(courseMap)

	// ------------------------------------------
	// map 的 key 是具有唯一性的
	// 怎么去检查 key 的唯一性
	// 必须要对比 key
	// == 或 !=
	// testMap := map[[]int]int{}

	// slice1 := []int{
	// 	1, 2, 3,
	// }
	//
	// slice2 := []int{
	// 	1, 2, 3,
	// }
	//
	// // invalid operation: slice1 == slice2 (slice can only be compared to nil)
	// fmt.Println(slice1 == slice2)

	// ----------------------------------------------
	// map1 := map[string]int{
	// 	"a": 1,
	// }
	//
	// map2 := map[string]int{
	// 	"a": 1,
	// }
	//
	// // invalid operation: map1 == map2 (map can only be compared to nil)
	// fmt.Println(map1 == map2)

	// --------------------------------------------------
	// arr1 := [3]int{1, 2, 3}
	// arr2 := [3]int{4, 5, 6}
	// arr3 := [3]int{7, 8, 9}
	//
	// testMap := map[[3]int]int{
	// 	arr1: 1,
	// 	arr2: 2,
	// 	arr3: 3,
	// }
	//
	// fmt.Println(testMap)

	// -------------------------------------------------------
	/**
	  map 的多类型的 value 问题
	*/
	// studentMap := map[string]interface{}{
	// 	"name":    "张三",
	// 	"age":     25,
	// 	"married": false,
	// }
	//
	// fmt.Println(studentMap)

	// ---------------------------------------------------------
	/**
	  map 的遍历的问题
	*/

	// courseMap := make(map[string]string, 3)
	//
	// courseMap["go"] = "Golang课程"
	// courseMap["java"] = "Java WEB课程"
	// courseMap["rust"] = "Rust课程"
	//
	// // map 是无序的集合，不要视图用 map 定义时的 key 顺序来描述程序对 key 的排序
	// for key, value := range courseMap {
	// 	fmt.Println(key, value)
	// }
	//
	// for key := range courseMap {
	// 	fmt.Println(key)
	// }
	//
	// for _, value := range courseMap {
	// 	fmt.Println(value)
	// }

	// ------------------------------------------------------------
	/**
	  map 对 key 的取值
	*/

	// courseMap := make(map[string]string, 3)
	//
	// courseMap["go"] = "Golang课程"
	// courseMap["java"] = "Java WEB课程"
	// courseMap["rust"] = "Rust课程"
	// courseMap["c"] = "C系列课程"
	// courseMap["cpp"] = ""

	// fmt.Println(courseMap["java"])
	// fmt.Println(courseMap["c"])

	// value, existed := courseMap["c"]
	// if existed {
	// 	fmt.Println(value)
	// }
	//
	// value2, existed2 := courseMap["cpp"]
	// fmt.Println(value2, existed2)
	//
	// if value, existed := courseMap["rust"]; existed {
	// 	fmt.Println(value, existed)
	// }

	/**
	  map key 的搜删除
	*/

	// courseMap := make(map[string]string, 3)
	//
	// courseMap["go"] = "Golang课程"
	// courseMap["java"] = "Java WEB课程"
	// courseMap["rust"] = "Rust课程"
	// courseMap["c"] = "C系列课程"
	// courseMap["cpp"] = ""
	//
	// delete(courseMap, "cpp")
	// fmt.Println(courseMap)
	//
	// // 删除一个 map 中没有定义的 key 不会报错
	// delete(courseMap, "js")
	// fmt.Println(courseMap)
}
