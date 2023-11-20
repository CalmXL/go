package main

/**
  Go 的并发哲学
	不要通过共享内存进行通信 (互斥锁、读写锁)
	要通过通信的方式来共享内存

	锁机制：用两个 G 去修改同一个变量值，变量值空间就是两个 G 的共享内存，
				我们需要通过锁来限制 G 的并发操作的原子化 (Go 不推荐的方式)

	通道机制: G 把数据推进一个通道内，让其他 G 通过这个通道去获取数据
				在这个通道中设计有发送与接收的规则
*/

func main() {
	// 声明一个 Channel
	// var c chan int
	// fmt.Println(c) // => <nil>

	// 声明创建一个 Channel
	// c := make(chan int)
	//
	// go func() {
	// 	c <- 1 // 发送或者推入  通道
	// 	c <- 2
	// 	c <- 3
	// }()
	//
	// v1 := <-c
	// v2 := <-c
	// v3 := <-c
	// fmt.Println(v1, v2, v3) // 1

	// ----------------------------------------
	// 无缓冲区
	// c := make(chan int)
	// 有缓冲区
	// c := make(chan int, 5)
	// s := []int{1, 2, 3, 4, 5}
	//
	// go func() {
	// 	/**
	// 	  close 环形结构的数组
	// 	  如果 Channel 永远不关闭，就会造成死锁的错误
	// 		GMP在死锁状态下，会将所有的 G 设置为休眠状态
	// 	*/
	// 	defer close(c)
	// 	for _, v := range s {
	// 		fmt.Printf("Sending %d\r\n", v)
	// 		c <- v
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()
	// time.Sleep(5 * time.Second)
	// for v := range c {
	// 	fmt.Printf("Received: %d\r\n", v)
	// }
	// TODO 多个值

	// ---------------------------------------------
	// CLOSE
	// c := make(chan int)
	//
	// go func() {
	// 	// Channel 被关闭后，还可以接收到一次数据
	// 	// 但是不能发送数据
	// 	defer close(c)
	// 	c <- 1
	// }()
	//
	// /**
	//   1. 发送的数据
	// 	ok -> true: 接收到发送的数据
	// */
	// v, ok := <-c
	// /**
	//   0: chan int => int 零值
	// 	ok -> false: 不是接收到发送的数据
	//
	// 	可以通过 ok 判断是不是接收的数据
	// */
	// v2, ok2 := <-c
	// fmt.Println(v, ok)
	// fmt.Println(v2, ok2)

	// ---------------------------------------------
	// 单向 Channel
	/**
	  本 Channel 只能发送或者只能接收
	*/

	// c := make(chan<- int)
	// c <- 1
	// // 无效运算: <- c (从仅发送类型 chan<- int 接收)
	// <-c

	// c := make(<-chan int)
	// // 无效运算: c <- 1 (发送到仅接收类型 <-chan int)
	// c <- 1

	/**
	  总结：这样的做法是完全封闭了 Channel 的发送或者接收
					业务需求在大部分情况下是不满足的
	*/
	// -----------------------------------------
	// var wg sync.WaitGroup
	// wg.Add(2)
	//
	// // 双向 Channel
	// c := make(chan int, 5)
	//
	// go func(c chan<- int) {
	// 	// 发送方关闭 Channel
	// 	defer close(c)
	// 	defer wg.Done()
	//
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Printf("Sending %d\r\n", i)
	// 		c <- i
	// 	}
	// }(c)
	//
	// go func(c <-chan int) {
	// 	defer wg.Done()
	// 	for v := range c {
	// 		fmt.Printf("Received %d\r\n", v)
	// 	}
	// }(c)
	//
	// time.Sleep(time.Second)

	// ------------------------------------------
	/**
	  死锁:

		锁: 阻塞
	  死锁: 阻塞必然无法解开的情况
				 或者是程序锁定了,
	*/

	// 情况1:
	/**
	  当无缓冲区 C 没有发送方，直接接收 => 死锁
		没有发送方的 Goroutine => 死锁
	*/
	// c := make(chan int)
	//
	// // 是一种阻塞
	// // fatal error: all goroutines are asleep - deadlock!
	// <-c

	// 情况2
	/**
	  无缓冲区的 C 没有接收方 => 死锁
		没有接收方的 Goroutine，C 的数据没有 G 消费，这就是一种无法解开的阻塞

		有缓冲区的 C 没有接收方 => 不会死锁
		因为 C 里有缓冲区数组保存这些数据， 所以不会死锁
	*/
	// c := make(chan int, 3)
	// c <- 1

	// 情况3：
	/**
		React
	    Provider 提供方 存在
	    Consumer 消费方 可以不存在

		Go:
			缓冲区满, 并继续推入数据到 Channel, 且没有接收方
	*/
	// c := make(chan int, 2)
	// c <- 1
	// c <- 2
	// c <- 3 // 阻塞

	// 情况4
	/**
	  c => nil => 没有缓冲区和接收方
		通道还没有创建，数据无法推入
	*/
	// var c chan int
	// // fatal error: all goroutines are asleep - deadlock!
	// c <- 1 // 阻塞

	// 情况5：
	// var wg sync.WaitGroup
	//
	// wg.Add(2)
	//
	// c := make(chan int, 5)
	//
	// go func(c chan<- int) {
	// 	defer wg.Done()
	//
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Printf("Sending %d \r\n", i)
	// 		c <- i
	// 	}
	// }(c)
	//
	// time.Sleep(3 * time.Second)
	//
	// go func(c <-chan int) {
	// 	defer wg.Done()
	//
	// 	for v := range c {
	// 		fmt.Printf("Receievd %d \r\n", v)
	// 	}
	// }(c)
	//
	// wg.Wait()

	// ------------------------------------------
	/**
	  两个slice 存放学生与老师的信息,
		通过 Channel 来传输这些信息，不同的 Channel 传递对应的信息
	  接收方将对应对应的信息放入对应 slice
	*/
	// 	var wg sync.WaitGroup
	// 	type Student struct {
	// 		id   int
	// 		name string
	// 		age  int
	// 		pass bool
	// 	}
	//
	// 	type Teacher struct {
	// 		id      int
	// 		name    string
	// 		age     int
	// 		subject string
	// 	}
	//
	// 	stSlice := make([]Student, 0)
	// 	thSlice := make([]Teacher, 0)
	//
	// 	stChan := make(chan Student)
	// 	thChan := make(chan Teacher)
	//
	// 	wg.Add(1)
	//
	// 	go func() {
	//
	// 		defer close(stChan)
	// 		defer close(thChan)
	// 		defer wg.Done()
	//
	// 		stChan <- Student{
	// 			id:   1,
	// 			name: "张三",
	// 			age:  18,
	// 			pass: false,
	// 		}
	//
	// 		thChan <- Teacher{
	// 			id:      1,
	// 			name:    "li Teacher",
	// 			age:     34,
	// 			subject: "数学",
	// 		}
	//
	// 		stChan <- Student{
	// 			id:   1,
	// 			name: "李四",
	// 			age:  18,
	// 			pass: false,
	// 		}
	//
	// 		thChan <- Teacher{
	// 			id:      1,
	// 			name:    "wang Teacher",
	// 			age:     44,
	// 			subject: "英语",
	// 		}
	// 	}()
	//
	// loop:
	// 	for {
	// 		select {
	// 		// 有数据可以接收
	// 		// 如果多个 chan 都有数据可以接收，则随机进入到某个 case 中
	// 		case st, ok := <-stChan:
	// 			if !ok {
	// 				break loop
	// 			}
	// 			stSlice = append(stSlice, st)
	// 		case th, ok := <-thChan:
	// 			if !ok {
	// 				break loop
	// 			}
	// 			thSlice = append(thSlice, th)
	// 		}
	//
	// 	}
	//
	// 	time.Sleep(5 * time.Second)
	//
	// 	for _, v := range stSlice {
	// 		fmt.Println(v)
	// 	}
	//
	// 	for _, v := range thSlice {
	// 		fmt.Println(v)
	// 	}
	//
	// 	wg.Wait()
}
