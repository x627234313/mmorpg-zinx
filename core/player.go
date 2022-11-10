package core

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/aceld/zinx/ziface"
	"github.com/x627234313/mmorpg-zinx/pb"
	"google.golang.org/protobuf/proto"
)

// 玩家对象
type Player struct {
	PId  uint32             // 玩家ID
	Conn ziface.IConnection // 玩家连接
	X    float32            // 平面的x轴坐标
	Y    float32            // 高度
	Z    float32            // 平面的y轴坐标
	V    float32            // 0-360度的角度
}

// 玩家ID 生成器
var PIdGen uint32 = 1
var IdLock sync.Mutex

// 创建一个玩家对象
func NewPlayer(conn ziface.IConnection) *Player {
	IdLock.Lock()
	pId := PIdGen
	PIdGen++
	IdLock.Unlock()

	player := &Player{
		PId:  pId,
		Conn: conn,
		X:    float32(160 + rand.Intn(20)),
		Y:    0,
		Z:    float32(120 + rand.Intn(10)),
		V:    0,
	}

	return player
}

// 提供一个给客户端发送消息的方法
// 主要是把pb的protobuf 数据序列化后，调用zinx的SendMsg方法
func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	// 将proto message 结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("proto marshal error:", err)
		return
	}

	// 将二进制数据通过 zinx框架的 SendMsg方法发送给客户端
	if p.Conn == nil {
		fmt.Println("Player connection is nil")
		return
	}

	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("Player send msg error:", err)
		return
	}

}

// 组建msgID=1 的 proto msg，同步player ID
func (p *Player) SyncPid() {
	proto_msg := &pb.SyncPid{
		Pid: int32(p.PId),
	}

	// 发送给客户端
	p.SendMsg(1, proto_msg)
}

// 组建msgID=200 的 proto msg，同步player 初始位置
func (p *Player) BroadCastStartPosition() {
	proto_msg := &pb.BroadCast{
		Pid: int32(p.PId),
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	// 发送给客户端
	p.SendMsg(200, proto_msg)
}
