package main

import (
	"container/list"
	"fmt"
)

func main() {
	/**
	  数组与链表都是有序列表

		数组：需要开辟一个连续的空间
				 查询是采用顺序位偏移的方式，性能效率高
				 插入，特别是中间插入，需要对后续元素进行移动，性能效率较低

		链表：是不需要开辟连续的空间，注入一个元素，开辟一个空间
				 插入只需要将前后的指针进行重新指向，不需要对其他元素进行存储移动 性能相对较高
				 查询，需要遍历元素进行查询，性能效率相对较低

		list 是 Go 中的双向链表
		type List struct {
			root Element // sentinel list element, only &root, root.prev, and root.next are used
			len  int     // current list length excluding (this) sentinel element
		}

		type Element struct {
			next, prev *Element
			list *List
			Value interface{}
		}
	*/

	// var linkList list.List
	// // var linkList = new(list.List).Init()
	// // var linkList = list.New()
	//
	// linkList.PushBack("Golang")
	// linkList.PushFront("iava")
	// linkList.PushBack("Rust")
	//
	// fmt.Println(linkList)
	/**
	  {{0x140000601e0 0x14000060210 <nil> <nil>} 3}
			Element                                  len
		{ prev, next, list Value}
	*/
	// fmt.Println(linkList.Len()) // 长度 => 3
	//
	// fmt.Println(linkList.Front()) // => &{0x140000741b0 0x14000074180 0x14000074180 iava}
	//
	// for i := linkList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }
	//
	// // 倒序遍历
	// for i := linkList.Back(); i != nil; i = i.Prev() {
	// 	fmt.Println(i.Value)
	// }

	// ------------------------------------------------------
	// 初始化一个链表 （返回一个初始化完毕的链表）
	// courseList := list.New()

	// 返回一个新增的元素(Element)
	// 填入的参数是一个 Value,而 PushBack 会帮助你包装成 Element 对象
	// newElement := courseList.PushBack("Golang")
	// fmt.Println(newElement)
	// fmt.Println(newElement.Value) // => Golang

	// goElement := courseList.PushBack("Golang")
	// javaElement := courseList.PushBack("Java")
	// rustElement := courseList.PushBack("Rust")
	//
	// fmt.Println(goElement, javaElement, rustElement)

	// courseList.InsertBefore("C++", javaElement)
	// courseList.InsertAfter("C", javaElement)

	// ------------------------------------------------

	// i := courseList.Front()
	//
	// for ; i != nil; i = i.Next() {
	// 	if i.Value.(string) == "Java" {
	// 		break
	// 	}
	// }
	//
	// courseList.InsertBefore("C++", i)
	// courseList.InsertAfter("C", i)
	//
	// for i := courseList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	// ------------------------------------------------
	// courseList := list.New()
	//
	// goElement := courseList.PushBack("Golang")
	// javaElement := courseList.PushBack("Java")
	// rustElement := courseList.PushBack("Rust")
	// fmt.Println(goElement, javaElement, rustElement)
	//
	// movingElement := courseList.Front()
	// targetElement := courseList.Front()

	// for i := courseList.Front(); i != nil; i = i.Next() {
	// 	if i.Value.(string) == "Java" {
	// 		targetElement = i
	// 	}
	//
	// 	if i.Value.(string) == "Rust" {
	// 		movingElement = i
	// 	}
	// }
	//
	// courseList.MoveBefore(movingElement, targetElement)
	// courseList.MoveAfter(movingElement, targetElement)

	// for i := courseList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	// -------------------------------------------------------

	// courseList := list.New()
	//
	// goElement := courseList.PushBack("Golang")
	// javaElement := courseList.PushBack("Java")
	// rustElement := courseList.PushBack("Rust")
	// fmt.Println(goElement, javaElement, rustElement)
	//
	// i := courseList.Front()
	// for ; i != nil; i = i.Next() {
	// 	if i.Value.(string) == "Rust" {
	// 		break
	// 	}
	// }
	//
	// courseList.MoveToFront(i)
	// courseList.MoveToBack(i)
	//
	// for i := courseList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	// -------------------------------------------------
	// 合并链表
	// courseList := list.New()
	//
	// goElement := courseList.PushBack("Golang")
	// javaElement := courseList.PushBack("Java")
	// rustElement := courseList.PushBack("Rust")
	// fmt.Println(goElement, javaElement, rustElement)
	//
	// studentList := list.New()
	//
	// studentList.PushBack("张三")
	// studentList.PushBack("李四")
	// studentList.PushBack("王五")
	//
	// // courseList.PushBackList(studentList)
	// courseList.PushFrontList(studentList)
	//
	// for i := courseList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	// -----------------------------------------------
	// 删除链表
	courseList := list.New()

	courseList.PushBack("Golang")
	courseList.PushBack("Java")
	courseList.PushBack("Rust")

	i := courseList.Front()

	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "Rust" {
			break
		}
	}

	courseList.Remove(i)

	for i := courseList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
