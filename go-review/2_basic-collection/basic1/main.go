package main

import "fmt"

func main() {
	// var classOne [3]string
	// var classThree [3]string
	//
	// classOne[0] = "1"
	// classOne[1] = "2"
	// classOne[2] = "3"
	//
	// classOne[1] = "200"
	// fmt.Println(classOne)
	//
	// fmt.Printf("%T  %T\r\n", classOne, classThree)
	
	// fmt.Println(classOne == classThree) // false
	//
	// classOne = classThree
	// fmt.Println(classOne) // [ ]
	//
	// fmt.Println(classOne == classThree) // true
	
	// for i := 0; i < len(classOne); i++ {
	// 	fmt.Println(classOne[i])
	// }
	
	// var nums [3]int
	// var bools [3]bool
	// fmt.Println(nums)
	// fmt.Println(bools)
	
	classOne := [3]string{1: "徐雷", 2: "徐红"}
	fmt.Println(classOne)
	
	for index, el := range classOne {
		fmt.Println(index, el)
	}
}
