package main

/*
切片 slice
*/

// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//
// var slice0 = arr[2:8]
// var slice1 = arr[:6]
// var slice2 = arr[5:10]
// var slice3 = arr[:]
// var slice4 = arr[:len(arr)-1]

func main() {

	// // 1. 声明切片
	// var s1 []int
	// if s1 == nil {
	// 	fmt.Println("s1 是空")
	// } else {
	// 	fmt.Println("不是空")
	// }
	//
	// // 2. :=
	// s2 := []int{}
	//
	// // 3. make
	// var s3 []int = make([]int, 0)
	// fmt.Println(s1, s2, s3) // => [] [] []
	//
	// // 4. 初始化赋值
	// var s4 []int = make([]int, 1, 10)
	// fmt.Println(s4)
	//
	// s5 := []int{3, 4, 5}
	// fmt.Println(s5)
	//
	// arr := [5]int{1, 2, 3, 4, 5}
	// var s6 []int
	//
	// // 左闭又开
	// s6 = arr[1:4]
	// fmt.Println(s6) // => [2, 3, 4]

	// 	*************************************************
	// fmt.Println(slice0)
	// fmt.Println(slice1)
	// fmt.Println(slice2)
	// fmt.Println(slice3)
	// fmt.Println(slice4)

	// **************************************************
	// make 创建切片: var slice []type = make([]type, len)
	//  slice := make([]type, len)
	//  slice := make([]type, len, cap)

	// 读写操作实际目标是底层数组
	// data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	// s := data[2:4]
	// s[0] = 100
	// s[1] += 100
	//
	// fmt.Println(s)    // [100 103]
	// fmt.Println(data) // [0 1 100 103 4 5 6 7 8]

	// 直接创建 slice, 可自动分配底层数组
	// s1 := []int{0, 1, 2, 3, 8: 100} // 使用索引号
	// fmt.Println(s1, len(s1), cap(s1))
	//
	// s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	// fmt.Println(s2, len(s2), cap(s2))
	//
	// s3 := make([]int, 6) // 省略 cap， cap = len
	// fmt.Println(s3, len(s3), cap(s3))

	// 使用 make 动态创建数组。还可以使用指针直接访问底层数组，退化成普通数组操作。
	// s := []int{0, 1, 2, 3}
	// s[2] += 100
	// // p := &s[2]
	// // *p += 100
	// fmt.Println(s)

	// [][]T
	// data := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{100, 200},
	// 	[]int{11, 22, 33},
	// }
	//
	// fmt.Println(data)

	// *******************************
	// d := [5]struct {
	// 	x int
	// }{}
	//
	// s := d[:]
	//
	// d[1].x = 100
	// s[0].x = 200
	//
	// fmt.Println(d)

	// ************************************
	// append
	// var a = []int{1, 2, 3}
	// fmt.Printf("slice a: %v\n", a)
	//
	// var b = []int{4, 5, 6}
	// fmt.Printf("slice b: %v\n", b)
	//
	// c := append(a, b...)
	// fmt.Printf("slice c: %v\n", c)
	//
	// d := append(c, 9)
	// fmt.Printf("slice d: %v\n", d)
	//
	// e := append(d, 10, 11, 21)
	// fmt.Printf("slice e: %v\n", e)

	// 超出 slice.cap 限制，就会重新分配底层数组， 即便原数组并未填满
	// s[low:high:max] 从切片 s 的索引位置 low 到 high 处所获得的切片
	// len = high - low, cap = max - low
	// data := [...]int{0, 1, 2, 3, 4, 10: 0}
	// s := data[:2:3]
	// fmt.Println(&s[0], &data[0])
	// // len 2, cap = 3
	// s = append(s, 100, 200)      // 一次 append 两个值，超出 s.cap 限制
	// fmt.Println(s, data)         // 重新分配底层数组，与数组无关
	// fmt.Println(&s[0], &data[0]) // 比起对底层数组起始指针

	// **********************************************
	// slice 中 cap 重新分配规律

	// s := make([]int, 0, 1)
	// c := cap(s)
	//
	// for i := 0; i < 50; i++ {
	// 	s = append(s, i)
	// 	if n := cap(s); n > c {
	// 		fmt.Printf("cap: %d -> %d\n", c, n)
	// 		c = n
	// 	}
	// }

	// cap: 1 -> 2
	// cap: 2 -> 4
	// cap: 4 -> 8
	// cap: 8 -> 16
	// cap: 16 -> 32
	// cap: 32 -> 64

	// ************************************
	// copy: 函数 copy 在两个 slice 间复制数据
}
