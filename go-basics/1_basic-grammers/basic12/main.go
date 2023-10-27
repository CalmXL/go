package main

import "fmt"

func main() {
	// 内置方法: len() 获取字符串字节长度
	// 英文字母不论大小写，每个字母占一个字节
	// 中文字母，每个字占 3 个字节

	// a := "I Love You!"
	// length := len(a)
	// fmt.Println(length) // => 11

	// a := "我爱你"
	// length := len(a)
	// fmt.Println(length) // => 9

	// 字母串是由字符组成的
	// [ I \s l o v e \s y o u ! ]
	// 字母或者空格 => byte 或者 rune
	// a := "I Love you!"
	// length := len(a)
	// fmt.Println(length)
	//
	// // 方括号代表的是切片
	// // 我要把 a 字符串转换成每个元素 byte 类型
	// aBytes := []byte(a)
	// // 每个元素都是一个 ASCII 码
	// fmt.Println(aBytes)      // [73 32 76 111 118 101 32 121 111 117 33]
	// fmt.Println(len(aBytes)) // => 11

	// a := "我爱你"
	// // byte => uint8 => 0 ~ 255
	// // 中文需要使用 rune
	// aBytes := []byte(a)
	// aRunes := []rune(a)
	// fmt.Println(aBytes) // => [230 136 145 231 136 177 228 189 160]
	// fmt.Println(aRunes) // => [25105 29233 20320]
	//
	// fmt.Println(len(aBytes)) // => 9
	// fmt.Println(len(aRunes)) // => 3

	// -----------------------------------------------
	// 转义符
	// 转义: 将符号意义转换为普通文本的字符特性
	// title1 := "他说: \"我最喜欢《悲惨世界》中的冉阿让\""
	// title2 := `他说: "我最喜欢《悲惨世界》中的冉阿让"`
	//
	// fmt.Println(title1, title2)

	/*
		\n  换行
		\r  回车
		\t Tab 键制表符
	*/

	// title3 := "我是\\小野森森\\。\r\n我今年37岁。\t欢迎大家听我的 GO 课程。"
	// fmt.Println(title3)

	// 制表符
	// fmt.Println("一\t二\t三\t四\t五\t六\t日")
	// fmt.Println("1\t2\t3\t4\t5\t6\t7")
	// fmt.Println("8\t9\t10\t11\t12\t13\t14")
	// fmt.Println("15\t16\t17\t18\t19\t20\t21")

	// -------------------------------------------------------
	a := "Hello"
	b := "Hello"
	c := "world"
	d := "Hola"
	e := "Helle"

	// 相等与不等的成立条件是字符串顺序与每个字母都相等
	res := a == b
	res1 := a != c
	fmt.Println(res, res1) // => true true

	// 比较大小时，字符串依次进行 ASCII 码对比
	res2 := a > c
	fmt.Println(res2) // => false

	res3 := a > b
	fmt.Println(res3) // => false

	res4 := a < d
	fmt.Println(res4) // => true

	res5 := a > e
	fmt.Println(res5) // => true
}
