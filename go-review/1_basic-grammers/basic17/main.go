package main

import "fmt"

func main() {
	name := "xulei-徐雷"
	nameRune := []rune(name)
	
	for index, element := range name {
		fmt.Println(index, element)
		// %c Unicode 转字符
		fmt.Printf("%d %c\r\n", index, element)
	}
	
	for index, _ := range nameRune {
		fmt.Printf("%d %c\r\n", index, nameRune[index])
	}
}
