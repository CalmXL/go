package main

import "fmt"

func main() {

	// type Test struct {
	// 	a int8  // 1
	// 	b int32 // 4
	// 	c bool  // 1
	// }

	// type Test struct {
	// 	a int8  // 1
	// 	c bool  // 1
	// 	b int32 // 4
	// }
	//
	// test := Test{
	// 	a: 1,
	// 	b: 2,
	// 	c: true,
	// }
	// // 12 -> 8
	// fmt.Println(unsafe.Sizeof(test)) // => 12

	// ------------------------------------------
	// type Todo struct {
	// 	id        int
	// 	content   string
	// 	completed bool
	// }
	//
	// var todo Todo
	//
	// todo.id = 1
	// todo.content = "content"
	// todo.completed = true

	// -------------------------------------------
	// type Address struct {
	// 	province string
	// 	city     string
	// 	district string
	// 	detail   string
	// }
	//
	// type LogisticsInfo struct {
	// 	orderId    string
	// 	productName string
	// 	address    struct {
	// 		province string
	// 		city     string
	// 		district string
	// 		detail   string
	// 	}
	// }

	// type LogisticsInfo struct {
	// 	orderId     string
	// 	productName string
	// 	province    string
	// 	address     Address
	// }
	//
	// myAddress := Address{
	// 	province: "黑龙江",
	// 	city:     "齐齐哈尔",
	// 	district: "讷河市",
	// 	detail:   "龙河镇",
	// }

	// // 匿名结构体
	// myAddress := struct {
	// 	province string
	// 	city     string
	// 	district string
	// 	detail   string
	// }{
	// 	province: "黑龙江",
	// 	city:     "齐齐哈尔",
	// 	district: "讷河市",
	// 	detail:   "龙河镇",
	// }
	//
	// fmt.Println(myAddress)

	// logisticsInfo := LogisticsInfo{
	// 	orderId:     "1234314",
	// 	productName: "HUAWEI Mate60 pro",
	// 	province:    "北京",
	// 	address: Address{
	// 		province: "黑龙江",
	// 		city:     "齐齐哈尔",
	// 		district: "讷河市",
	// 		detail:   "龙河镇",
	// 	},
	// }
	// logisticsInfo.address.district = "黑河"
	// fmt.Println(logisticsInfo)

	// ---------------------------------------------
	type Address struct {
		province string
		city     string
		district string
		detail   string
	}

	type LogisticsInfo struct {
		orderId     string
		productName string
		province    string
		Address
	}

	logisticsInfo := LogisticsInfo{
		"1234314",
		"HUAWEI Mate60 pro",
		"北京",
		Address{
			province: "黑龙江",
			city:     "齐齐哈尔",
			district: "讷河市",
			detail:   "龙河镇",
		},
	}
	logisticsInfo.province = "上海市"
	fmt.Println(logisticsInfo.province) // => 访问的是外部的
}
