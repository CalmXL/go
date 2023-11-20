package main

import "fmt"

func main() {
	// resInt := plusInt(1, 2)
	// resFloat := plusFloat(1.2, 1.3)
	// resString := plusString("a", "b")
	//
	// fmt.Println(resInt, resFloat, resString)

	/*
		  在函数调用的时候，[] 中可以填入具体的类型来限制参数传递的类型
										如果不写中括号，就有Go通过函数的参数与返回值推断出最后的结果的类型
	*/
	// [int] => 类型实参
	resInt := plus[int](1, 2)
	resFloat := plus(1.1, 2.3)
	resString := plus("a", "b")

	fmt.Println(resInt, resFloat, resString)

	// numSlice := []int{1, 2, 3}
	zhangsanMap := map[string]interface{}{
		"name": "张三",
		"age":  18,
	}

	lisiMap := map[string]interface{}{
		"name": "李四",
		"age":  23,
	}

	studentSlice := make([]map[string]interface{}, 2)
	studentSlice = append(studentSlice, zhangsanMap, lisiMap)

	forEachPrintSlice(studentSlice)

	// fmt.Println(equals(1, 1))
	// fmt.Println(equals("a", "a"))
	// fmt.Println(equals(zhangsanMap, lisiMap))

	// type CommonSlice[T int | string] []T
	// var intSlice CommonSlice[int] = []int{}
	//
	// type CommonMap[K int | string, V int | string] map[K]V

	var sumSlice SumSlice[int] = []int{1, 2, 3, 4, 5}

	res := sumSlice.Sum()
	fmt.Println(res)
}

/*
	 提出问题
		同样的参数接口，同样的单值返回，同样函数逻辑
		我们需要定义三个函数来完成 这个功能是不是合适

		遇到的问题
			三个函数因为参数返回值的类型不同
		如果我们要整合成一个函数？
		plus(a ?, b ?) ? {}
		泛型：泛指一个类型（不是一个具体的类型）
			泛型是一个因为类型不确定而需要占位的一个标识符
			这个标识符可以理解为类似于函数形参变量的东西
*/

/*
  T 实际上就是一个泛型，泛型的作用就是在函数定义时先占位一个不确定的类型的位置
  函数调用的时候，通过参数类型推断来确定实际泛型对应的具体数据位置

	在 Go 语言中，泛型是需要在函数名后面进行定义才能使用的
  定义泛型的方法
		1.在函数名后面跟[]
    2. 在中括号里写入泛型标识[T]
		3. 在 Go 语言中，泛型定义必须有类型约束(告诉泛型定义，这个泛型到底有多少可选的具体类型)
		4. 泛型标识原理上是可以写任何字母或者单词
				T => Type
				E => Element
				K => Key
				V => Value
*/

// 泛型函数
// T 就是类型的形参
func plus[T int | string | float64](a T, b T) T {
	// T 没有得到类型约束，就会导致，有可能加法运算无效
	return a + b
}

func plusInt(a, b int) int {
	return a + b
}

func plusFloat(a, b float64) float64 {
	return a + b
}

func plusString(a, b string) string {
	return a + b
}

// -----------------------------------------------------
// func forEachPrintSlice[E string | int](s []E) {
// 	for _, e := range s {
// 		fmt.Println(e)
// 	}
// }

// any 代表任何类型都可以接收
func forEachPrintSlice[E any](s []E) {
	for _, e := range s {
		fmt.Println(e)
	}
}

func equals[T comparable](a, b T) bool {
	return a == b
}

type SumSlice[T int | float64] []T

func (cs SumSlice[T]) Sum() T {
	var res T
	for _, e := range cs {
		res += e
	}
	return res
}
