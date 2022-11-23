package apis

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/x627234313/mmorpg-zinx/core"
	"github.com/x627234313/mmorpg-zinx/pb"
	"google.golang.org/protobuf/proto"
)

// 世界聊天的路由业务
type WorldChatApi struct {
	znet.BaseRouter
}

func (wc *WorldChatApi) Handle(request ziface.IRequest) {
	// 解析客户端传递进来的proto协议
	proto_msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("talk proto unmarshal error:", err)
		return
	}

	// 当前的聊天信息是哪个玩家发送的
	pid, err := request.GetConnection().GetProperty("pid")

	// 根据pid得到对应的玩家对象
	player := core.WorldMgr.GetPlayer(pid.(uint32))

	// 将消息广播给全部在线的玩家
	player.Talk(proto_msg.Content)
}
