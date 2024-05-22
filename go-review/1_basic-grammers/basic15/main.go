package main

import "fmt"

func main() {
	math := 70
	chemistry := 80
	physics := 90
	biology := 100
	
	subject := ""
	
	fmt.Println("请输入学科名: ")
	fmt.Scanln(&subject)
	
	switch subject {
	case "数学":
		fmt.Printf("数学成绩: %d", math)
		break
		fmt.Println(111)
		fallthrough
	case "化学":
		fmt.Printf("化学成绩: %d", chemistry)
	case "物理":
		fmt.Printf("物理成绩: %d", physics)
	case "生物":
		fmt.Printf("生物成绩: %d", biology)
	}
}
