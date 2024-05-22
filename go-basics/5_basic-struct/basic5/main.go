package main

import "fmt"

/*
	泛型结构体
*/

type S[T int | float64] struct {
	a T
	b T
}

/*
	todoList->todo {
		id: int,
		content: []string | string
		completed: bool
	}
*/
type Todo[T []string | string] struct {
	id        int
	content   T
	completed bool
}

/**
slice 的增删改查封装
*/

type SliceData[T any] struct {
	data []T
	size int
}

func (sd SliceData[T]) Size() int {
	return sd.size
}

func (sd *SliceData[T]) Push(elements ...T) (index int) {
	sd.data = append(sd.data, elements...)
	sd.size += len(elements)
	index = len(sd.data) - 1
	return index
}

func (sd *SliceData[T]) Remove(index int) (element T) {
	element = sd.data[index]
	sd.data = append(sd.data[:index], sd.data[index+1:]...)
	sd.size -= 1
	return element
}

func (sd *SliceData[T]) Set(index int, element T) (oldElement T, newElement T) {
	oldElement = sd.data[index]
	newElement = element
	sd.data[index] = newElement
	return oldElement, newElement
}

func (sd *SliceData[T]) GetSlice() []T {
	return sd.data
}

func (sd *SliceData[T]) Get(index int) T {
	return sd.data[index]
}

/*
data.ForEach(func (el, index, slice) {
})
*/
func (sd *SliceData[T]) ForEach(cb func(
	element T,
	index int,
	slice []T,
)) {
	for index, el := range sd.data {
		cb(el, index, sd.data)
	}
}

func (sd *SliceData[T]) Map(cb func(
	element T,
	index int,
	slice []T,
) T) []T {
	var newSlice []T = make([]T, len(sd.data))

	for index, el := range sd.data {
		newEl := cb(el, index, sd.data)
		newSlice[index] = newEl
	}
	return newSlice
}

func main() {
	//s := S[int]{
	//	a: 1,
	//	b: 2,
	//}
	//
	//res := s.a + s.b
	//fmt.Println(res)
	//
	//s2 := S[float64]{
	//	a: 1.1,
	//	b: 2.2,
	//}
	//
	//res2 := s2.a + s2.b
	//fmt.Println(res2)

	// ---------------------------------------
	//todoList1 := make([]Todo[string], 0)
	//
	//todoList1 = append(todoList1, Todo[string]{
	//	id:        1,
	//	content:   "todo",
	//	completed: false,
	//})
	//
	//todoList2 := make([]Todo[[]string], 0)
	//
	//todoList2 = append(todoList2, Todo[[]string]{
	//	id:        1,
	//	content:   []string{"a", "b"},
	//	completed: false,
	//})
	//
	//fmt.Println(todoList1, todoList2)

	// ------------------------------------------------------
	ds := SliceData[int]{
		data: []int{1, 2, 3, 4, 5},
		size: 5,
	}

	fmt.Println(ds.GetSlice())

	index := ds.Push(6, 7, 8, 9)
	fmt.Println(index, ds.GetSlice())

	oldEl, newEl := ds.Set(5, 600)
	fmt.Println(oldEl, newEl)

	delEl := ds.Remove(3)
	fmt.Println(delEl, ds.GetSlice(), ds.Size())

	el := ds.Get(7)
	fmt.Println(el)

	ds.ForEach(func(el int, index int, slice []int) {
		fmt.Println(el, index, slice)
	})

	newSlice := ds.Map(func(el int, index int, slice []int) int {
		el += 1
		return el
	})

	fmt.Println(newSlice)
}
