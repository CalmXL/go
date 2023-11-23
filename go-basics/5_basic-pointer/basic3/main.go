package main

import "fmt"

/*
	  nil:
			1. Go 语言中的零值|空值 => 引用类型的空值
			2. nil 不是一种数据类型，也不是关键字，Go 语言标识符

		Go 语言中数据的零值
			bool: false
			number: 0
			string: ""

			pointer: nil
			slice: nil
			map: nil
			channel: nil
			func: nil
			interface: nil

			struct没有零值
				struct 是数据结构，内部有结构的装载具体类型
				struct 内部的字段有可能有零值
*/
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

	// ----------------------------------------
	// nil 不是一种数据类型
	// var a *int = nil
	// fmt.Printf("%T", a)
	// // => *int
	// fmt.Printf("%T", nil)
	// // => <nil> 没有类型 unTyped

	/*
		如果没有显式类型，则无法分配 nil
		无法将 'nil' 用作类型 Type

		1. 这样的写法是一种需要 Go 进行类型推断的写法
		2. 赋值 nil, nil 不是一种数据类型，所以没办法进行准确推断的类型判断
		3. nil 作为某些类型的零值，不止一种，该推断成哪一种零值，不确定
		4. 赋值 nil 的俩种条件:
			a. 显示定义变量的类型
			b. 类型必须符合 nil 类该类型零值的条件
	*/
	// a := nil
	// var a int = nil
	// var a []int = nil

	// ---------------------------------------------------------
	// test := func() int {
	// 	return nil
	// }

	// test := func () []int {
	// 	return nil
	// }

	// ---------------------------------------------------------
	// nil 不是 Go 语言的关键字(不可以使用)
	// nil := 1
	// fmt.Println(nil) // => 1

	// ---------------------------------------------------------
	/*
		  引用在 Go 中是不能做比较的
			那么相等性判断也是不可做的
	*/
	// var s1 []int = nil
	// // var s2 []int = nil
	//
	// // 无效运算: s1 == s2 (在 []int 中未定义运算符 ==)
	// // if s1 == s2 {
	// //
	// // }
	//
	// if s1 == nil {
	// 	fmt.Println("nil")
	// } else {
	// 	fmt.Println("not nil")
	// }

	// s := make([]int, 0) // 空间为 0 的 slice
	// // panic: runtime error: index out of range [0] with length 0
	// // s[0] = 1
	// fmt.Println(s) // => []
	//
	// a := 1
	// // 无法将 'nil' 转换为类型 'int'
	// if a == nil {
	//
	// }
	type Adress struct {
		province string
		city     string
	}

	type Student struct {
		name string
		age  int
		// hobbies []int
		// hobbies map[string]string
		address Adress
	}

	s1 := Student{
		name: "张三",
		age:  18,
		address: Adress{
			province: "黑龙江",
			city:     "齐齐哈尔",
		},
	}

	s2 := Student{
		name: "张三",
		age:  18,
		address: Adress{
			province: "黑龙江",
			city:     "齐齐哈尔",
		},
	}

	fmt.Printf("%p\r\n", &s1)
	fmt.Printf("%p\r\n", &s2)

	if s1 == s2 {
		fmt.Println("ture")
	} else {
		fmt.Println("false")
	}

	// ------------
	// var m1 map[int]int
	// // panic: assignment to entry in nil map
	// // m1[1] = 1
	//
	// fmt.Printf("%p", &m1)

	// ------------------------------
	// empty  map
	var m = make(map[int]int, 0)

	m[0] = 0
}
