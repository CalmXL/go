package main

import "fmt"

/*
*

	  条件判断

		if 条件 => 程序1
		if 条件 => 程序2
		else   => 程序3

		单条件判断：
			if 条件 => 程序

		多条件判断
			if 条件      => 程序1
			else if 条件 => 程序2
			else if 条件 => 程序3
			else        => 程序4


		Go 语言
			条件：
				1. 推荐不打括号
				2. 条件必须 布尔表达式
*/
func main() {
	// x := 1 // 1 0

	// ❌ 非布尔值 'x' (类型 int) 用作条件
	// 判断条件必须是布尔表达式
	// if x {
	// 	fmt.Println("Success")
	// } else {
	// 	fmt.Println("error")
	// }

	// math := 60
	// physics := 80
	// chemistry := 90
	// biology := 55

	/**
	  1. 数学必须是 60 分以上
		2. 物理化学生物至少有一个是 60 分以上
			physics >= 60 || chemistry >= 60 || biology >= 60

		满足两个条件则通过，不满足则不通过

		为什么 Go 语言 if 判断条件不用括号括起来
			如果复杂的条件判断，要用括号扩起来
			如果外层也有括号，就显得混乱，且没有必要性
	*/

	// if (math >= 60) && (physics >= 60 || chemistry >= 60 || biology >= 60) {
	// 	fmt.Println("pass")
	// } else {
	// 	fmt.Println("Not Pass")
	// }

	/**
	  如果满足条件是唯一的，那么就需要多条件判断
		如果有多个条件都有可能满足，那么就需要单独写 if 判断
	*/
	// if (math >= 60) && (physics >= 80) {
	// 	fmt.Println("Physics is excellent")
	// } else if (math >= 60) && (chemistry >= 80) {
	// 	fmt.Println("Chemistry is excellent")
	// } else if (math >= 60) && (biology >= 80) {
	// 	fmt.Println("Biology is excellent")
	// } else {
	// 	fmt.Println("No excellent subject")
	// }

	// -----------------------------------

	// if (math >= 60) && (physics >= 80) {
	// 	fmt.Println("Physics is excellent")
	// }
	//
	// if (math >= 60) && (chemistry >= 80) {
	// 	fmt.Println("Chemistry is excellent")
	// }
	//
	// if (math >= 60) && (biology >= 80) {
	// 	fmt.Println("Biology is excellent")
	// }

	// --------------------------------------
	// if math >= 60 {
	// 	if physics >= 80 {
	// 		fmt.Println("Physics is excellent")
	// 	}
	//
	// 	if chemistry >= 80 {
	// 		fmt.Println("Chemistry is excellent")
	// 	}
	//
	// 	if biology >= 80 {
	// 		fmt.Println("Biology is excellent")
	// 	}
	// }

	// -------------------------------------------
	// x := 1
	// fmt.Scanln(&x)
	// fmt.Println(x)

	// year := 0
	// month := 0
	//
	// fmt.Println("请输入年份")
	// _, err := fmt.Scanln(&year)
	// if err != nil {
	// 	return
	// }
	//
	// fmt.Println("请输入月份")
	// _, err = fmt.Scanln(&month)
	// if err != nil {
	// 	return
	// }
	//
	// if month == 2 {
	// 	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
	// 		fmt.Println("天数： 29")
	// 	} else {
	// 		fmt.Println("天数: 28")
	// 	}
	// } else {
	// 	if month <= 12 {
	// 		if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
	// 			fmt.Println("天数: 30")
	// 		} else {
	// 			fmt.Println("天数: 31")
	// 		}
	// 	} else {
	// 		fmt.Println("请输入正确月份")
	// 	}
	// }

	/**
	  1. 判断是否登录
			1. 如果登录可以查询数学、物理、化学、生物分数
			2. 如果没登录，则要求用户正确输入用户名密码登录
	*/
	logged := false
	isLogin(logged)
}

func verify(username string, password string) bool {
	if (username == "xulei") && (password == "123456") {
		return true
	}
	return false
}

func isLogin(logged bool) {
	subject := ""
	username := ""
	password := ""
	if logged {
		fmt.Println("登录成功。")
		fmt.Println("请输入你要查询的科目分数?")
		fmt.Scanln(&subject)

		res := querySubject(subject)

		fmt.Printf("当前科目的分数为 %d", res)
	} else {
		fmt.Println("请输入你的用户名: ")
		fmt.Scanln(&username)
		fmt.Println("请输入你的密码: ")
		fmt.Scanln(&password)

		logged = verify(username, password)
		isLogin(logged)
	}
}

func querySubject(subject string) int {
	math := 60
	physics := 80
	chemistry := 90
	biology := 55

	if subject == "math" {
		fmt.Println(math)
		return math
	}

	if subject == "physics" {
		return physics
	}

	if subject == "chemistry" {
		return chemistry
	}

	if subject == "biology" {
		return biology
	}

	return 0
}
