package main

import "fmt"

// 条件分支 switch
/**
  if 常用在条件是范围的情况下
		a > 10 && a <= 20
		常用在函数内部，进行函数程序返回
		if err != nil {
			return
		}
		Println(123)

	switch: 常用在条件是一个一个的固定的值的时候
		if a == 1 {
		} else if a == 2 {
		} else if a == 3 {
		} else {}

		switch a {
				case 1:
				case 2:
				case 3:
				default:
		}

	switch 是不需要 break 的，
	格式：
		switch 变量 {
			case 值1:
				程序1
			case 值2:
				程序3
			case 值3:
				程序3
			default:
				程序4
		}
*/
func main() {
	math := 70
	// chemistry := 80
	// physics := 90
	// biology := 100
	//
	// subject := ""
	//
	// fmt.Println("请输入学科名称: ")
	// _, err := fmt.Scanln(&subject)
	// if err != nil {
	// 	return
	// }

	/**
	  1. switch 不会主动穿透 case, 不需要 case 内部写 break, 方式穿透 case
		2. break 会强制让程序跳出 switch，本 case 中 break 以下的代码都是无法访问的代码
		3. fallthrough 穿透 case, 可以穿透到下一个 case
	*/
	// switch subject {
	// case "数学", "shuxue", "math", "mathematics":
	// 	fmt.Printf("数学成绩 %d", math)
	// 	// fallthrough
	// case "化学":
	// 	fmt.Printf("化学成绩 %d", chemistry)
	// 	break
	// 	fmt.Println("123")
	// case "物理":
	// 	fmt.Printf("物理成绩 %d", physics)
	// case "生物":
	// 	fmt.Printf("生物成绩 %d", biology)
	// default:
	// 	fmt.Println("输入的学科名称不存在")
	// }

	switch {
	case math < 60:
		fmt.Println("F")
	case math >= 60 && math < 70:
		fmt.Println("D")
	case math >= 70 && math < 80:
		fmt.Println("C")
	case math >= 80 && math < 90:
		fmt.Println("B")
	case math >= 90 && math < 100:
		fmt.Println("A")
	default:
		fmt.Println("This is no-valid")
	}
}
