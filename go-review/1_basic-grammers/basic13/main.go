package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("abc")
	// fmt.Println("def")
	//
	// fmt.Print("abc")
	// fmt.Println("def")
	//
	// name := "xulei"
	// age := 18
	// fmt.Println("My name is " + name + ", I am " + strconv.Itoa(age) + " years old.\r\n")
	// fmt.Printf("My name is %s, I am %d years old.\r\n", name, age)
	//
	// res := fmt.Sprintf("My name is %s, I am %d years old.\r\r", name, age)
	// fmt.Println(res)
	//
	// fmt.Printf("Age 类型：%T\r\n", age)
	//
	// str := "abc"
	// strRunes := []rune(str)
	// fmt.Printf("%v", strRunes)
	
	// strings
	// name := "xulei"
	// age := 18
	// var stringBuilder strings.Builder
	//
	// stringBuilder.WriteString("My name is ")
	// stringBuilder.WriteString(name)
	// stringBuilder.WriteString(". ")
	// stringBuilder.WriteString("I am ")
	// stringBuilder.WriteString(strconv.Itoa(age))
	// stringBuilder.WriteString(" years old.\r\n")
	//
	// res := stringBuilder.String()
	// fmt.Println(res)
	
	// strings.Contain
	// title := "我正在学习 Go 语言，以后将成为优秀的 Go 语言开发者。"
	// res1 := strings.Contains(title, "Go")
	// res2 := strings.Contains(title, "Java")
	// fmt.Println(res1, res2) // => true false
	
	// strings.Count
	// res3 := strings.Count(title, "Go")
	// fmt.Println(res3) // => 2
	
	// str1 := "ABCD-EFGH-HIJK-LMNO"
	// res4 := strings.Split(str1, "-")
	// fmt.Println(res4)
	//
	// for i := 0; i < len(res4); i++ {
	// 	fmt.Println(res4[i])
	// }
	
	// str2 := `"This is Golang"`
	// res5 := strings.Index(str2, "Go")
	// fmt.Println(res5)
	
	// str3 := "Go developer, Go engineer"
	// res6 := strings.Replace(str3, "Go", "Java", 1)
	// fmt.Println(res6)
	
	// str5 := " @ Go developer # "
	// res7 := strings.Trim(str5, "@# ")
	// fmt.Println(res7)
	//
	// str6 := "GO DEVELOPER"
	// res8 := strings.ToLower(str6)
	// fmt.Println(res8)
	//
	// res9 := strings.ToUpper(res8)
	// fmt.Println(res9)
	
	str10 := "$$ Go developer **"
	res10 := strings.HasPrefix(str10, "$")
	res11 := strings.HasSuffix(str10, "*")
	
	fmt.Println(res10, res11)
}
