package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func test() (int, error) {
	// return 1, fmt.Errorf("出现错误,错误码为 %d", 500)
	return 1, errors.New("出错")
}

func readFile() {
	// defer catchPainc()
	file, err := os.Open("./_basic-function/basic5/desc.txt")
	
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("关闭文件失败")
		}
	}(file)
	
	if err != nil {
		fmt.Println("error:", err)
		panic("打开文件失败")
	}
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(34, line)
	}
}

func catchPainc() {
	if err := recover();
		err != nil {
		fmt.Println("catch panic", err)
	}
}

func main() {
	// res, err := test()
	// fmt.Println(res, err)
	//
	// defer fmt.Println("defer1")
	// defer fmt.Println("defer2")
	// defer fmt.Println("defer3")
	//
	// fmt.Println("main1")
	// fmt.Println("main2")
	// fmt.Println("main3")
	
	readFile()
}
