1. Channel 拿来做什么？

   Channel 通信 => 通知 => 告诉对方要干什么？

   TodoList需要干什么？
        1. 数据的增删改查
            -> 发送信号 -> 2

        2. 显示操作后的结果，显示操作目录
            Print 2 <- 接收信号和数据


2. 需要完成什么任务？
    todo => {} => struct
    todoList => [] | channel
         增删改查的接口

    chan <- 增删改查 =》 Type
            数据 => Payload  Todo  [] Todo