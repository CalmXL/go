package main

import (
	"fmt"
	"strconv"
)

/*
*

	  3 行两列数据
		"前端班", "60人", "2000结课", "终身"
		"Java班", "55人", "2500结课", "终身"
		"Golang班", "45人", "1800结课", "终身"

			[
				["前端班", "60人", "2000结课", "终身"],
				["Java班", "55人", "2500结课", "终身"],
				["Golang班", "45人", "1800结课", "终身"]
			]
*/
func main() {
	// [行数][列数]
	// var classInfo [3][4]string
	//
	// classInfo[0] = [4]string{"前端班", "60人", "2000结课", "终身"}
	// classInfo[1] = [4]string{"Java班", "55人", "2500结课", "终身"}
	// classInfo[2] = [4]string{"Golang班", "45人", "1800结课", "终身"}

	// fmt.Println(classInfo)

	// 遍历行
	// for i := 0; i < len(classInfo); i++ {
	// 	// classInfo[0]
	// 	for j := 0; j < len(classInfo[i]); j++ {
	// 		fmt.Print(classInfo[i][j] + " ")
	// 	}
	// 	fmt.Println()
	// }

	// for _, rowEl := range classInfo {
	// 	for _, colEl := range rowEl {
	// 		fmt.Print(colEl + " ")
	// 	}
	// 	fmt.Println()
	// }

	// ------------------------------

	// 杨辉三角
	var triangleArr [5][]int

	for i := 0; i < len(triangleArr); i++ {
		triangleArr[i] = make([]int, i+1)
		triangleArr[i][0] = 1
		triangleArr[i][i] = 1

		for j := 1; j < len(triangleArr[i])-1; j++ {
			triangleArr[i][j] = triangleArr[i-1][j-1] + triangleArr[i-1][j]
		}
	}

	for i := 0; i < len(triangleArr); i++ {
		for j := 0; j < ((5 - (i + 1)) * 2); j++ {

			fmt.Print("*")
		}

		for k := 0; k < len(triangleArr[i]); k++ {
			fmt.Print(strconv.Itoa(triangleArr[i][k]) + "  ")
		}

		fmt.Println()
	}

	// fmt.Println(triangleArr)

}
