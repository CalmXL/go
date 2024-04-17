package main

import "fmt"

func main() {
	// type I interface {
	// }
	//
	// var b bool
	// var n int
	// var s string
	// var p *int
	// var sl []int
	// var m map[string]string
	// var c chan int
	// var f func()
	// var i I
	//
	// fmt.Println(b, n, s, p, sl, m, c, f, i)
	
	// fmt.Printf("%T", nil)
	
	// use of untyped nil in assignment
	// a := nil
	// cannot use nil as int value in variable declaration
	// var a int = nil
	// var a []int = nil
	
	// fmt.Println(a)
	
	// nil := 1
	// fmt.Println(nil)
	
	// var s1 []int = nil
	// var s2 []int = nil
	//
	// // 无效运算: s1 == s2 (在 []int 中未定义运算符 ==)
	// if s1 == s2 {
	//
	// }
	
	// s := make([]int, 0)
	// // panic: runtime error: index out of range [0] with length 0
	// s[0] = 1
	// fmt.Println(s)
	
	// type Address struct {
	// 	province string
	// 	city     string
	// }
	//
	// type Student struct {
	// 	name    string
	// 	age     int
	// 	address Address
	// }
	//
	// s1 := Student{
	// 	name: "张三",
	// 	age:  18,
	// 	address: Address{
	// 		province: "黑龙江",
	// 		city:     "齐齐哈尔",
	// 	},
	// }
	//
	// s2 := Student{
	// 	name: "张三",
	// 	age:  18,
	// 	address: Address{
	// 		province: "黑龙江",
	// 		city:     "齐齐哈尔",
	// 	},
	// }
	//
	// fmt.Printf("%p\r\n", &s1)
	// fmt.Printf("%p\r\n", &s2)
	// // 0x140000a6040
	// // 0x140000a6080
	//
	// if s1 == s2 {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }
	
	// var m1 map[int]int
	// // panic: assignment to entry in nil map
	// m1[1] = 1
	
	// empty map
	var m = make(map[int]int, 2)
	
	m[0] = 0
	m[1] = 1
	m[2] = 2
	m[3] = 3
	
	fmt.Println(m)
}
