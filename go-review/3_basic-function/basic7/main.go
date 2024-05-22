package main

import "fmt"

func main() {
	
	// resInt := plusInt(1, 2)
	// resFloat := plusFloat(1.1, 2.2)
	// resString := plusString("a", "b")
	//
	// fmt.Println(resInt)
	// fmt.Println(resFloat)
	// fmt.Println(resString)
	
	// resInt := plus[int](1, 2)
	// fmt.Println(resInt)
	//
	// resFloat := plus[float64](1.1, 2.2)
	// fmt.Println(resFloat)
	//
	// resString := plus[string]("a", "b")
	// fmt.Println(resString)
	
	type CommonMap[K int | string, V int | string] map[K]V
	
	var m2 CommonMap[string, string] = map[string]string{
		"name": "xulei",
		"age":  "18",
	}
	
	fmt.Println(m2)
	
	// var m1 = map[int]string{
	// 	1: "2",
	// 	2: "3",
	// }
	//
	// fmt.Println(m1)
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

func plus[T int | string | float64](a, b T) T {
	return a + b
}
