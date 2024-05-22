package main

import "fmt"

//
// type S[T int | float64] struct {
// 	a T
// 	b T
// }

// type Todo[T []string | string] struct {
// 	id        int
// 	content   T
// 	completed bool
// }

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
	return
}

func (sd *SliceData[T]) Remove(index int) (element T) {
	element = sd.data[index]
	sd.data = append(sd.data[0:index], sd.data[index+1:]...)
	sd.size -= 1
	return element
}

func (sd *SliceData[T]) Set(index int, element T) (oldElement T, newElement T) {
	oldElement = sd.data[index]
	newElement = element
	sd.data[index] = newElement
	return oldElement, newElement
}

func (sd *SliceData[T]) GetSlice() []T { return sd.data }

func (sd *SliceData[T]) Get(index int) (element T) {
	return sd.data[index]
}

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
	
	// s := S[int]{
	// 	a: 1,
	// 	b: 2,
	// }
	//
	// res := s.a + s.b
	// fmt.Println(res) // => 3
	//
	// s2 := S[float64]{
	// 	a: 1.1,
	// 	b: 2.2,
	// }
	//
	// res2 := s2.a + s2.b
	// fmt.Println(res2)
	
	sliceData := SliceData[int]{
		data: []int{1, 2, 3, 4, 5},
		size: 5,
	}
	
	fmt.Println(sliceData.Size())
	
	sliceData.Push(11, 22, 33)
	fmt.Println(sliceData)
	
	removeItem := sliceData.Remove(4)
	fmt.Println(removeItem)
	fmt.Println(sliceData)
	
	oldEl, newEl := sliceData.Set(0, 100)
	fmt.Println(oldEl, newEl)
	fmt.Println(sliceData)
	
	val := sliceData.Get(0)
	fmt.Println(val)
	
	s := sliceData.GetSlice()
	fmt.Println(s)
	
	sliceData.ForEach(func(el int, index int, slice []int) {
		fmt.Println(el, index, slice)
	})
	
	newSlice := sliceData.Map(func(el int, index int, slice []int) int {
		newEl := el + 100
		return newEl
	})
	
	fmt.Println(newSlice)
}
