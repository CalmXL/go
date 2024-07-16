### 1. DNS 解析

正常的网页访问流程机制
1. URL 地址的填写： https://www.baidu.com
2. URL -> IP 地址： => DNS服务器 => xxx.xxx.xxx.xxx
3. 向其的服务器程序发送请求
4. 服务器接收请求并处理
5. 服务器向客户端响应资源

另一种访问
1. URL 地址的填写： https://www.baidu.com
2. URL -> IP 地址： => hosts文件 => xxx.xxx.xxx.xxx
3. 向其的服务器程序发送请求
4. 服务器接收请求并处理
5. 服务器向客户端响应资源

### 2. 微服务中心(发现服务)
    统一管理服务配置、监控服务的健康状况、提供服务的发现、接入
    服务的注册、注销、服务的信息、
    HTTP访问、DNS解析
    
### 3.Consul

    一个服务发现的工具 + 可视化管理服务的工具
    Zookepper
    Eureka
    etcd
    Consul

### 4.Docker
    exe => ? => mac 不兼容
    .java => .class => windows x 
    => 环境的支持性配置

    1. VM => Virtual Machine 虚拟机
        exe => mac => window 支持的 VM
        .java/.class => windows => JVM

        问题: 
            1. 资源占用比较多
            2. 系统总是有一些权限问题和访问操作的，配置比较繁杂
            3. 毕竟要启动一个完整的操作系统，启动速度就无法保证绝对快

    2. 基于 linux 操作系统开发了一套虚拟化技术 => Linux 容器
         它不是操作系统，基于宿主系统的一个进程隔离

        1. 启动比较快
        2. 资源占用优化
        3. 体积小，不是一个系统的体量

    3. Docker, 对 linux 容器这样的技术进行功能性封装的产物 容器化工具
 