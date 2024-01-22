package main

import "fmt"

func main() {
	// stuSlice := []string{"张三", "李四", "王五", "赵六", "刘能", "小沈阳"}
	
	// fmt.Println(stuSlice[1])
	
	// slice1 := stuSlice[1:4]
	// fmt.Println(slice1) // => ["李四", "王五", "赵六"]
	//
	// slice2 := stuSlice[:]
	// fmt.Println(slice2) // => [ "张三", "李四", "王五", "赵六", "刘能", "小沈阳"]
	//
	// slice3 := stuSlice[1:]
	// fmt.Println(slice3) // => [李四 王五 赵六 刘能 小沈阳]
	//
	// slice4 := stuSlice[:4]
	// fmt.Println(slice4) // => [张三 李四 王五 赵六]
	//
	// slice5 := stuSlice[:10]
	// // ❌ panic: runtime error: slice bounds out of range [:10] with capacity 6
	// fmt.Println(slice5)
	
	// 遍历
	// for i := 0; i < len(stuSlice); i++ {
	// 	fmt.Println(stuSlice[i])
	// }
	//
	// for index, el := range stuSlice {
	// 	fmt.Println(index, el)
	// }
	
	// 添加元素
	// slice := []string{"张三", "李四"}
	// stuSlice := make([]string, 3)
	//
	// stuSlice[0] = "小明"
	// stuSlice[1] = "小红"
	// stuSlice[2] = "小明"
	//
	// slice = append(slice, stuSlice...)
	// fmt.Println(slice)
	//
	// slice2 := append(slice, stuSlice[1:]...)
	// fmt.Println(slice2)
	
	// // 删除元素
	// stuSlice := []string{"张三", "李四", "王五", "赵六", "刘能", "小沈阳"}
	//
	// // 删除第一个元素
	// slice1 := stuSlice[1:]
	// fmt.Println(slice1)
	//
	// // 删除最后一个元素
	// slice2 := stuSlice[:len(stuSlice)-1]
	// fmt.Println(slice2)
	//
	// // 删除中间元素
	// slice3 := append(stuSlice[:2], stuSlice[4:]...)
	// fmt.Println(slice3)
	
	// 复制、拷贝
	// stuSlice := []string{"张三", "李四", "王五", "赵六", "刘能", "小沈阳"}
	//
	// // 两个切片指向同一个切片引用
	// newSlice := stuSlice
	//
	// newSlice[0] = "xulei"
	//
	// fmt.Println(stuSlice) // => [xulei 李四 王五 赵六 刘能 小沈阳]
	//
	// stuSlice[1] = "wang"
	//
	// fmt.Println(newSlice) // => [xulei wang 王五 赵六 刘能 小沈阳]
	
	// 拷贝
	stuSlice := []string{"张三", "李四", "王五", "赵六", "刘能", "小沈阳"}
	newSlice := make([]string, len(stuSlice))
	
	// copy 返回切元素的长度
	res := copy(newSlice, stuSlice)
	
	fmt.Println(res, stuSlice, newSlice)
	
	newSlice[1] = "迪丽热巴"
	
	fmt.Println(stuSlice, newSlice)
	// 查看指针
	fmt.Printf("%p %p", &stuSlice, &newSlice)
}
