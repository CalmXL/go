package main

/*
service => 是一个完成某个特定功能的封装

		=> 跟 function 很像
		=> function 功能是将一段特定程序（可复用）封装在同一个作用域下

	  1. 从语言编程角度上来说， function 实际上是一个提供可执行的程序
		2. 从项目开发的角度上来说，function 实际上是提供完成某个特定的任务的程序

	  3. service 一个完成某个项目所需要完成的任务的程序体
				Micro-Service => 将一个特定的功能从主程序拆分出来单独开发的一个程序
				Micro 是相对于所有完成各个不同功能的程序集合的开发方式来说的
*/
func main() {
	// 可被函数封装的一种编程状态
	// a := 1
	// b := 2
	//
	// c := plus(a, b)
	//
	// d := 3
	// e := 4
	// 函数 => 完成了一个加法的任务
	// 服务 => 完成了一个加法的服务
	//        服务对象是谁呢？ main 程序
	// LPC => Local Procedure Call

	/*
		在同一个进程下，或者同一个计算机多个进程下或多个线程下进行服务的调用的方式就是 => LPC

		LPC => 从属于 => IPC: Inter-Process Communication 进程间通信
	*/
	// f := plus(d, e)

	/*
			当你的 plus 方法被编写在另一个计算机上的话（云函数）RPC => 从属于 => IPC
				Remote Procedure Call 远程过程调用
				1. 寻址（找到特定服务器的特定的服务程序）
		    2. 找到要调用的函数
				3. 函数执行的结果会响应回客户端
	*/

	/*
		http:localhost:3000/data  RPC调用
		app.get('/data', (req, res) => {
			res.json(data)
		})
	*/

	/**
	  服务端的微服务程序
			微服务与微服务之间需要进行业务上的通讯
			但微服务是一个单独的服务器程序进行部署发布的
			所以不能像普通的本地调用函数的方式进行对方的微服务的功能调用
			微服务是通过 RPC 技术对另一个微服务中的功能进行调用
	*/
}

// func plus(a int, b int) int {
// 	return a + b
// }
