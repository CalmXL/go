package main

// type MyInt int
// type MyInt2 = int

func main() {
	// var a MyInt = 1
	// var b int = 2
	// var c MyInt2 = 3
	//
	// // 无效运算: a == b(类型 MyInt 和 int 不匹配)
	// // fmt.Println(a == b)
	//
	// res := int(a) + b
	// fmt.Println(res)
	//
	// res2 := b + c
	// fmt.Println(res2)
	
	// type Account float64
	//
	// var accountBalance Account
	// var income Account = 123.12
	// total := accountBalance + income
	//
	// fmt.Println(total) // => 123.12
	
	type StudentMap map[string]interface{}
	
	var studentMap StudentMap = map[string]string{}
}
