package main

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/x627234313/mmorpg-zinx/apis"
	"github.com/x627234313/mmorpg-zinx/core"
)

// 客户端上线后的hook函数
func OnConnStart(conn ziface.IConnection) {
	// 创建player
	player := core.NewPlayer(conn)

	// 发送 msgID:1 的消息
	player.SyncPid()

	// 发送 msgID:200 的消息
	player.BroadCastStartPosition()

	// 将当前上线玩家添加到世界管理模块中
	core.WorldMgr.AddPlayer(player)

	// 将当前连接绑定一个玩家ID 的属性
	conn.SetProperty("pid", player.PId)

	// ====同步周边玩家上线信息，显示周边玩家信息====
	player.SyncSurrounding()
}

func main() {
	// 创建zinx server 句柄
	s := znet.NewServer()

	// 连接创建、销毁的hook函数
	s.SetOnConnStart(OnConnStart)

	// 注册一些路由业务
	s.AddRouter(2, &apis.WorldChatApi{})

	// 启动服务
	s.Serve()
}
