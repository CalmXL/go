package main

import "fmt"

func compute(
	a int,
	b int,
	method string,
	cb func(
	a int,
	b int,
	sign string,
	res int,
) string,
) string {
	_sign := ""
	_res := 0
	
	switch method {
	case "PLUS":
		_sign = "+"
		_res = a + b
	case "MINUS":
		_sign = "-"
		_res = a + b
	case "MULTIPLY":
		_sign = "*"
		_res = a * b
	case "DIVISION":
		_sign = "/"
		_res = a / b
	default:
		panic("error method")
	}
	
	return cb(a, b, _sign, _res)
}

func main() {
	res := compute(1, 2, "DIVISION", func(
		a int,
		b int,
		sign string,
		res int,
	) string {
		result := fmt.Sprintf(`%d %s %d = %d`, a, sign, b, res)
		
		return result
	})
	
	fmt.Println(res)
}
