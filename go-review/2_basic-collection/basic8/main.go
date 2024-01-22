package main

import "fmt"

func main() {
	// var slice1 []string
	// var slice2 []string
	//
	// // ❌ invalid operation: slice1 == slice2 (slice can only be compared to nil)
	// fmt.Println(slice1 == slice2)
	
	// map1 := map[string]string{}
	// map2 := map[string]string{}
	//
	// // ❌ invalid operation: map1 == map2 (map can only be compared to nil)
	// fmt.Println(map1 == map2)
	
	// myInfo := map[string]string{}
	//
	// myInfo["name"] = "xiaoyesensen"
	// myInfo["age"] = "18"
	//
	// fmt.Println(myInfo)
	
	// myInfo := map[string]string{
	// 	"name": "xulei",
	// 	"age":  "18",
	// }
	//
	// fmt.Println(myInfo)
	
	// myInfo := make(map[string]string, 3)
	//
	// myInfo["name"] = "xulei"
	// myInfo["age"] = "18"
	//
	// fmt.Println(myInfo)
	//
	// myInfo["height"] = "177"
	// myInfo["weight"] = "78"
	//
	// fmt.Println(myInfo)
	
	// 重复 key
	// myInfo := map[string]string{
	// 	"name": "xulei",
	// 	"age":  "18",
	// 	// "name": "xuhong",
	// }
	//
	// myInfo["age"] = "20"
	// fmt.Println(myInfo)
	
	// intMap := map[int]string{}
	//
	// intMap[1] = "1"
	// intMap[2] = "2"
	//
	// fmt.Println(intMap)
	
	// boolMap := map[bool]int{}
	//
	// boolMap[true] = 1
	// boolMap[false] = 0
	//
	// fmt.Println(boolMap)
	
	// myMap := map[string][]int{}
	//
	// myMap["slice1"] = []int{1, 2, 3}
	// myMap["slice2"] = []int{2, 3, 4}
	//
	// fmt.Println(myMap)
	
	// myMap := make(map[string]int, 3)
	//
	// myMap["age"] = 18
	// myMap["height"] = 180
	// myMap["weight"] = 78
	//
	// fmt.Println(myMap["s"]) // 默认返回对应类型的 0 值
	//
	// _, exist := myMap["age"]
	// fmt.Println(exist) // => true
	//
	// _, exist2 := myMap["s"]
	// fmt.Println(exist2)
	
	// var mySlice []string
	
	// // => panic: assignment to entry in nil map
	// var myMap map[int]int
	//
	// myMap[0] = 1
	
	myMap := make(map[string]int, 3)
	
	myMap["age"] = 18
	myMap["height"] = 180
	myMap["weight"] = 78
	
	// delete 内置方法
	// delete(myMap, "age")
	//
	// // 删除不存在的 key, 不会有任何操作和报错
	// delete(myMap, "music")
	
	// fmt.Println(myMap)
	
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
