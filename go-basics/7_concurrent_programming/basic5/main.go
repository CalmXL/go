package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// var step = 1
var s []int
var c = make(chan int)

/**
   问题:
    1. 两个 Goroutine 是没有办法相互阻塞的
	2. 全局变量 step 的使用，形成了一个复杂的全局变量更新的问题
	3. 过度的关心外部的操作
*/

func addStep() {
	defer wg.Done()
	step := 1
	//for {
	//	step++
	//	time.Sleep(1 * time.Second)
	//}

	for {
		c <- step
		step++
		time.Sleep(1 * time.Second)
	}
}

func addEl() {
	wg.Done()
	//for {
	//	s = append(s, step)
	//	fmt.Println(s)
	//	time.Sleep(2 * time.Second)
	//}

	for {
		select {
		case step := <-c:
			s = append(s, step)
			fmt.Println(s)
		}

		time.Sleep(2 * time.Second)
	}
}

func main() {
	wg.Add(2)
	go addStep()
	go addEl()

	wg.Wait()
}

/**
  type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
*/
