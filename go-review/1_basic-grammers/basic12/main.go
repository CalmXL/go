package main

import "fmt"

func main() {
	// a := "我爱你"
	// aRunes := []rune(a)
	// aBytes := []byte(a)
	// fmt.Println(aRunes)
	// fmt.Println(aBytes)
	//
	// fmt.Println(len(aRunes))
	
	// title1 := "他说: \"我最喜欢《悲惨世界》中的冉阿让\""
	// title2 := `他说: "我最喜欢《悲惨世界》中的冉阿让"`
	//
	// fmt.Println(title1)
	// fmt.Println(title2)
	
	/**
	  \n 换行
		\r 回车
		\t Tab 制表符
	*/
	
	// title3 := "我是\\小野森森\\。\r\n我今年37岁。\t欢迎大家听我的 Go 课程。"
	// fmt.Println(title3)
	
	// fmt.Println("一\t二\t三\t四\t五\t六\t日")
	// fmt.Println("1\t2\t3\t4\t5\t6\t7")
	// fmt.Println("8\t9\t10\t11\t12\t13\t14")
	// fmt.Println("15\t16\t17\t18\t19\t20\t21")
	
	a := "Hello"
	b := "Hello"
	c := "world"
	// d := "Hola"
	// e := "Helle"
	
	res := a == b
	fmt.Println(res)
	res1 := a != c
	fmt.Println(res1)
	
}
