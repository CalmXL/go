package main

import "fmt"

func main() {
	
	var classInfo [3][4]string
	
	classInfo[0] = [4]string{"前端班", "60人", "2000结课", "终身"}
	classInfo[1] = [4]string{"Java班", "55人", "2500结课", "终身"}
	classInfo[2] = [4]string{"Golang班", "45人", "1800结课", "终身"}
	
	fmt.Println(classInfo)
	
	// 遍历行
	
	// for i := 0; i < len(classInfo); i++ {
	// 	for j := 0; j < len(classInfo[i]); j++ {
	// 		fmt.Println(classInfo[i][j] + "")
	// 	}
	// 	fmt.Println()
	// }
	
	for _, rowEl := range classInfo {
		for _, colEl := range rowEl {
			fmt.Print(colEl + "   ")
		}
		fmt.Println()
	}
}
