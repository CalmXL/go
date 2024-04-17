package main

import "fmt"

func main() {
	// type Test struct {
	// 	a int8  // 1
	// 	b int32 // 4
	// 	c bool  // 1
	// }
	//
	// type Test2 struct {
	// 	a int8  // 1
	// 	b bool  // 1
	// 	c int32 // 4
	// }
	//
	// fmt.Println(unsafe.Sizeof(Test{
	// 	a: 1,
	// 	b: 2,
	// 	c: true,
	// })) // => 12
	//
	// fmt.Println(unsafe.Sizeof(Test2{
	// 	a: 1,
	// 	b: true,
	// 	c: 2,
	// })) // => 8
	//
	// type Test3 struct {
	// 	a int64   // 8
	// 	b bool    // 1
	// 	c string  // 不确定的？？
	// 	d float64 // 8
	// }
	//
	// type Test4 struct {
	// 	a int64
	// 	b bool
	// 	d float64
	// 	c string
	// }
	//
	// fmt.Println(unsafe.Sizeof(Test3{
	// 	a: 1,        // 8
	// 	b: false,    // 4
	// 	c: "亲爱的", // 16
	// 	d: 1.1,      // 8
	// }))
	//
	// type t struct {
	// 	// a int64
	// 	c string
	// }
	
	type Address struct {
		province string
		city     string
		district string
		detail   string
	}
	
	type LogisticsInfo struct {
		orderId     string
		productName string
		address     Address
	}
	
	myAddress := Address{
		province: "黑龙江",
		city:     "齐齐哈尔",
		district: "讷河市",
		detail:   "龙河镇",
	}
	
	// 匿名结构体
	myAddress2 := struct {
		province string
		city     string
		district string
		detail   string
	}{
		province: "黑龙江",
		city:     "齐齐哈尔",
		district: "讷河市",
		detail:   "龙河镇",
	}
	
	fmt.Println(myAddress2)
	
	myLogisticsInfo := LogisticsInfo{
		orderId:     "942255012",
		productName: "R9000P",
		address:     myAddress,
	}
	
	fmt.Println(myLogisticsInfo)
}
