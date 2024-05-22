package main

import "fmt"

// 鸭子类型

type Duck interface {
	walk()
	shout()
}

type Bird struct {
	legs int
}

type Dog struct {
	legs int
}

func (bd *Bird) walk() {
	fmt.Println("Bird is walking")
}

func (bd *Bird) shout() {
	fmt.Println("Bird is shouting")
}

func (dg *Dog) walk() {
	fmt.Println("Dog is walking")
}

func (dg *Dog) shout() {
	fmt.Println("Dog is shouting")
}

func doSth(animal Duck) {
	animal.walk()
	animal.shout()
}

func main() {
	
	doSth(&Bird{
		legs: 2,
	})
	
	doSth(&Dog{
		legs: 4,
	})
}
