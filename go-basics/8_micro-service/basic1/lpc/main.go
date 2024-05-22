package main

import "fmt"

/**
  LPC Local Procedure Call 本地过程调用

	多线程程序调用
	多进程程序调用
*/

func main() {
	res := Compute("MINUS", 1, 2)
	fmt.Println(res)
}

func Compute(t string, a int, b int) int {
	switch t {
	case "PLUS":
		return plus(a, b)
	case "MINUS":
		return minus(a, b)
	default:
		panic("传递的 type 有误")
	}
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}
