package main

import (
	"fmt"
	"strings"
)

/**
格式化输出
	3种
		1. Println, Print
		2. Printf, Sprintf
		3. strings.Builder
*/

func main() {

	/**
	Println => Print line
		P => 包外使用的方法
		fmt => Println => 包外使用
		有回车换行 => \r\n
	*/

	// fmt.Println("abc")
	// fmt.Println("def")
	//
	// fmt.Print("abc")
	// fmt.Println("def")
	//
	// fmt.Println("-------------------------------------------------")
	// name := "xiaoyesensen"

	// age := 18

	// Printf => print format
	// 性能：Print > Printf
	// fmt.Printf("My name is " + name + ", I am " + strconv.Itoa(age) + " years old.\r\n")
	// fmt.Printf("My name is %s, I am %d years old.\r\n", name, age)
	//
	// // sPrintf => 将格式化后的字符串返回
	// res := fmt.Sprintf("My name is %s, I am %d years old.\r\n", name, age)
	// fmt.Println(res)

	// %T 查看类型
	// fmt.Printf("Age 类型：%T\r\n", age)
	//
	// str := "abc"
	// strRunes := []rune(str)
	// fmt.Printf("%#v", strRunes)
	//
	// fmt.Println("---------------------------------------------------")

	// strings => 字符串方法包
	// Builder => WriteString()
	// 性能相当高

	// var stringBuilder strings.Builder
	// stringBuilder.WriteString("My name is ")
	// stringBuilder.WriteString(name)
	// stringBuilder.WriteString(". ")
	// stringBuilder.WriteString("I am ")
	// stringBuilder.WriteString(strconv.Itoa(age))
	// stringBuilder.WriteString(" years old.\r\n")
	//
	// res2 := stringBuilder.String()
	// fmt.Println(res2)

	/**
	strings => 字符串方法包
		Contains 检查包含
	*/

	title := "我正在学习 Go 语言，以后将成为优秀的 Go 语言开发者。"

	/**
	Contains: 检查字符串是否包含子字符串
		props:
			1. 字符串
			2. 子字符串
		return boolean
	*/
	res1 := strings.Contains(title, "Go")
	res2 := strings.Contains(title, "Java")
	fmt.Println(res1, res2)

	fmt.Println("--------------------------------------------------")

	/**
	Count: 检查子字符串在字符串中出现的次数
		props:
			1. 字符串
			2. 子字符串
		return int
	*/
	res3 := strings.Count(title, "Go")
	fmt.Println(res3)

	fmt.Println("--------------------------------------------------")

	/**
	Split: 分割字符串组成切片 []
		props:
			1. 字符串
			2. 分割符
		return []切片
	*/
	str1 := "FDKS-FIDO-VDIS-GFOD-GCODW"
	res4 := strings.Split(str1, "-")
	fmt.Println(res4)

	for i := 0; i < len(res4); i++ {
		fmt.Println(res4[i])
	}

	fmt.Println("--------------------------------------------------")

	/**
	Index: 子字符串在字符串中的位置
		props:
			1. 字符串
			2. 子字符串
		return int
		index: 下标 有序集合中，每一个元素对应的从 0 开始顺序排序的位置
	*/

	str2 := `"This is Golang"`
	res5 := strings.Index(str2, "is")
	fmt.Println(res5)

	fmt.Println("--------------------------------------------------")

	/**
	Replace: 替换字符串中的某个子串
		props:
			1. 字符串
			2. 待替换的字符串
			3. 新的替换字符串
			4. 替换次数(-1 全部替换)
	*/
	str3 := "Go developer, Go engineer"
	res6 := strings.Replace(str3, "Go", "Java", 2)
	fmt.Println(res6)

	fmt.Println("--------------------------------------------------")

	/**
	Trim: 去掉字符串前后指定的字符
		props:
			1. 字符串
			2. 自定要去掉前后的 xx 字符
	*/
	str5 := " @ Go developer #  "
	res7 := strings.Trim(str5, " @#")
	fmt.Println(res7)

	fmt.Println("--------------------------------------------------")

	/**
	ToLower: 把所有字符转成小写
	ToUpper: 把所有字符转成大写
	*/
	str6 := "GO DEVELOPER"
	res8 := strings.ToLower(str6)
	fmt.Println(res8)

	res9 := strings.ToUpper(res8)
	fmt.Println(res9)

	/**
	HasPrefix 是否有某个前缀字符串
	HasSuffix 是否有某个后缀字符串
		props:
			1. 字符串
			2. 字符
	*/

	str10 := "$$GO DEVELOPER##"
	res10 := strings.HasPrefix(str10, "$$")
	res11 := strings.HasSuffix(str10, "##")
	fmt.Println(res10)
	fmt.Println(res11)
}
