package main

import "github.com/aceld/zinx/znet"

func main() {
	// 创建zinx server 句柄
	s := znet.NewServer()

	// 连接创建、销毁的hook函数

	// 注册一些路由业务

	// 启动服务
	s.Serve()
}
