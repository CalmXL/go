package main

import "fmt"

/*
	布尔类型: 是与非、真与假、对与错、成立与不成立

		1 > 2: 非、假、错、不成立 => false
    2 > 1: 是、真、对、成立 => true

	byte 类型：
		1. byte 类型是存放字符编码的类型 (uint8)
		2. 字符一定用单引号
		3. 字符存储本质上是字符编码的存储

		ASCII: American Standard Code information Interchange
					美国信息交换标准代码
		4. ASCII 本质上就是一个字符对应一个整数的字符与数字之间的交换
			 a => 97
		5. byte 字符的存储就是对字符对应的 ASCII 码进行存储

		rune 类型：
			1. rune 类型是存放字符编码的类型 (int32)
			2. 存储更多字符的编码，可以存储包含所有 unicode 码对应的字符
				 unicode 本质上就是对 ASCII 的拓展
			3. rune 存储的是字符对应的编码
			4. 而且字符一定要用单引号

	String 类型：
		1. 存储字符串类型
		2. 值要用双引号
		3. 在 GO 语言当中，一个中文字占三个字节

*/

func main() {
	var b1 bool
	b1 = true
	fmt.Println(b1) // => true

	var b2 = false
	fmt.Println(b2) // => false

	var b3 bool = true
	fmt.Println(b3)

	b4 := false
	fmt.Println(b4)

	fmt.Println("---------------------------------------")

	var a byte
	a = 'a'
	fmt.Printf("%c", a)
	fmt.Println(a)

	var b byte = 'b'
	fmt.Printf("%c", b)
	fmt.Println(b)

	var bb byte = 'b'
	fmt.Printf("%c", bb)
	fmt.Println(bb)

	var c = 'c'
	fmt.Printf("%c", c)
	fmt.Println(c)

	d := 'd'
	fmt.Printf("%c", d)
	fmt.Println(d)

	// var e byte = d + 1 x
	var bbb byte = bb + 1
	fmt.Printf("%c", bbb)
	fmt.Println(bbb)

	fmt.Println("---------------------------------------")

	var ch1 rune = '我'
	fmt.Printf("%c", ch1)
	fmt.Println(ch1)

	//var ch1 rune = 25105
	//fmt.Printf("%c", ch1)
	//fmt.Println(ch1)

	//var ch2 = '爱'
	//fmt.Printf("%c", ch2)
	//fmt.Println(ch2)

	var ch2 = 23323
	fmt.Printf("%c", ch2)
	fmt.Println(ch2)

	ch3 := '你'
	fmt.Printf("%c", ch3)
	fmt.Println(ch3)

	var ch4 rune = ch3 + 1
	fmt.Printf("%c", ch4)
	fmt.Println(ch4)

	fmt.Println("---------------------------------------")

	var title string
	title = "我爱你"
	fmt.Println(title)

	var title2 string = "我爱你"
	fmt.Println(title2)

	var title3 = "我爱你"
	fmt.Println(title3)

	title4 := "我爱你"
	fmt.Println(title4)

	title5 := "亲爱的"

	// 拼接字符串，用 +
	fullTitle := title5 + "," + title4
	fmt.Println(fullTitle)

	fullTitleLen := len(fullTitle)
	fmt.Println(fullTitleLen) // => 19 => 3(一个中文字 3 个字符) * 6 + 1

	fmt.Println("---------------------------------------")

	text1 := "I am a boy."
	text2 := "I am 18 years old."

	fullText := text1 + text2
	fmt.Println(fullText)
	fmt.Println(len(fullText))

	fmt.Println("---------------------------------------")

}
